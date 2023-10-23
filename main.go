package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"

	"htmx-test/api"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	db, err := sql.Open("postgres", "host=localhost user=root password=secret dbname=gourmeg_2 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	t, err := template.ParseGlob("public/views/*.html")

	if err != nil {
		log.Fatalf("unable to parse templates: %b", err)
	}

	e := echo.New()
	e.Renderer = &Templates{templates: t}

	e.Use(middleware.Logger())
	e.Static("/api", "public/api")
	e.Static("/css", "public/css")
	e.Static("/js", "public/js")

	e.GET("/", func(c echo.Context) error {
		lists, err := api.GetTopLevelLists(db)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%b", err))
		}
		return c.Render(http.StatusOK, "home", lists)
	})

	e.GET("/list/:id", func(c echo.Context) error {
		id_str := c.Param("id")
		id, err := strconv.ParseInt(id_str, 10, 0)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
		}

		l, err := api.GetList(db, int(id))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%b", err))
		}

		return c.Render(http.StatusOK, "list_page", l)
	})

	e.GET("/recipe/:id", func(c echo.Context) error {
		id_str := c.Param("id")
		id, err := strconv.ParseInt(id_str, 10, 0)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
		}

		recipe, err := api.GetRecipe(db, int(id))
		return c.Render(http.StatusOK, "recipe_page", recipe)
	})

	e.GET("/api/item/:id/edit", func(c echo.Context) error {
		id_str := c.Param("id")
		id, err := strconv.ParseInt(id_str, 10, 0)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
		}

		var i api.Item
		err = c.Bind(&i)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
		}

		i.Id = int(id)

		return c.Render(http.StatusOK, "edit_item", i)
	})

	e.GET("/api/item/:id", func(c echo.Context) error {
		id_str := c.Param("id")
		id, err := strconv.ParseInt(id_str, 10, 0)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
		}
		var i api.Item
		c.Bind(&i)

		i, err = api.GetItem(db, int(id), i.Type)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
		}
		return c.Render(http.StatusOK, "item", i)
	})

	e.PUT("/item/:id", func(c echo.Context) error {
		var i api.Item
    err := c.Bind(&i)
    fmt.Printf("%s\n", i.Type)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
		}

		recipe, err := api.PutItem(db, i)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
		}
		return c.Render(http.StatusOK, "item", recipe)
	})

	e.POST("/recipe", func(c echo.Context) error {
		r := new(api.Recipe)
		if err := c.Bind(r); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		err := api.PostRecipe(db, r)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
		}

		return c.Render(http.StatusOK, "add_recipe_result", r)
	})

	e.POST("/list_search", func(c echo.Context) error {
		list := c.FormValue("list")

		items, err := api.SearchList(db, list)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
		}

		return c.Render(http.StatusOK, "list_of_items", items)
	})

	e.GET("/add", func(c echo.Context) error {
		return c.Render(http.StatusOK, "add.html", nil)
	})

	e.Debug = true
	e.Logger.Fatal(e.Start(":1323"))
}
