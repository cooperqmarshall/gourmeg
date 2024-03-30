package api

import (
	"fmt"
	"htmx-test/db"
	"io"
	"net/http"
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

    html, err := fetch_recipe_html(r.Url)
    if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    _, err = extract_recipe_ldjson(html)
    _ = new(recipe.Recipe)
    // r.Read_jsonld()
    if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

	err = db.PostRecipe(handler.DB, r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.Render(http.StatusOK, "add_recipe_result", r)
}

func fetch_recipe_html(url string) (io.Reader, error) {
    res, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    return res.Body, nil
}

func extract_recipe_ldjson(r io.Reader) ([]byte, error) {
    t := html.NewTokenizer(r)

    for {
        if t.Next() == html.ErrorToken {
            return nil, t.Err()
        }
        fmt.Printf("%s", t.Token())
        return nil, nil
    }
}

func (handler Handler) GetRecipe(c echo.Context) error {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
	}

	recipe, err := db.GetRecipe(handler.DB, int(id))
	fmt.Printf("%d, %d", id, recipe.Id)
	return c.Render(http.StatusOK, "recipe_page", recipe)
}
