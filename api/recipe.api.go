package api

import (
	"database/sql"
	"fmt"
	"gourmeg-htmx/db"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cooperqmarshall/recipe"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/html"
)

func (handler Handler) PostRecipe(c echo.Context) error {
	r := new(db.Recipe)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if len(r.Url) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "url cannot be empty")
	}

	ignore_duplicates := c.QueryParam("ignore_duplicates")
	if ignore_duplicates != "true" {
		fmt.Printf("%s\n", ignore_duplicates)
		// check if recipe already added
		r2, err := db.GetRecipeFromURL(handler.DB, r.Url)
		if err != nil {
			if err == sql.ErrNoRows {
				// no matches found
			} else {
				return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error checking for recipe duplicates: \"%s\"", err))
			}
		}

		if r2.Id != 0 {
			// hacky way of passing current list to html template
			r2.ListId = r.ListId

			return c.Render(http.StatusOK, "duplicate_recipe_confirm", r2)
		}
	}

	err := fetch_recipe(r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = db.PostRecipe(handler.DB, r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.Render(http.StatusOK, "item", db.Item{Id: r.Id, Name: r.Name, Type: "recipe"})
}

func fetch_recipe(r *db.Recipe) error {
	html, err := fetch_recipe_html(r.Url)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	re, err := extract_recipe_ldjson(html)
	if err != nil {
		if err.Error() != "EOF" {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("Error when extracting recipe from website \"%s\"", err))
		}
	}

	if re.Name == "" {
		r.Name = r.Url
	} else {
		r.Name = re.Name
	}
	r.Ingredients = re.Ingredients
	r.Instructions = re.Instructions

	return nil
}

func fetch_recipe_html(u string) (io.Reader, error) {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, fmt.Errorf("invalid URL \"%s\"", u)
	}

	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func extract_recipe_ldjson(reader io.Reader) (recipe.Recipe, error) {
	t := html.NewTokenizer(reader)
	found_ldjson := false
	var r recipe.Recipe

	for {
		if t.Next() == html.ErrorToken {
			return r, t.Err()
		}
		token := t.Token()

		if found_ldjson {
			err := r.Read_jsonld([]byte(token.Data))
			if err != nil && (len(r.Ingredients) != 0 || len(r.Instructions) != 0) {
				return r, nil
			}
			found_ldjson = false
		}

		if len(token.String()) < 7 || token.String()[1:7] != "script" {
			continue
		}

		for _, attr := range token.Attr {
			if attr.Key == "type" && attr.Val == "application/ld+json" {
				found_ldjson = true
			}
		}
	}
}

func (handler Handler) GetRecipe(c echo.Context) error {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("id param (%s) not a number", id_str))
	}

	recipe, err := db.GetRecipe(handler.DB, int(id))
	return c.Render(http.StatusOK, "recipe_page", recipe)
}

func (handler Handler) RefetchRecipe(c echo.Context) error {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("id param (%s) not a number", id_str))
	}

	r, err := db.GetRecipe(handler.DB, int(id))

	err = fetch_recipe(r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = db.UpdateRecipe(handler.DB, r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("Error updating recipe: (%s)", err))
	}

	return c.Render(http.StatusOK, "recipe_page", r)
}
