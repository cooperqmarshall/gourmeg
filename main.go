package main

import (
	"database/sql"
	"html/template"
	"io"

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
	e := echo.New()
	e.Use(middleware.Logger())

	db, err := sql.Open("postgres", "host=localhost user=root password=secret dbname=gourmeg_2 sslmode=disable")
	if err != nil {
		e.Logger.Fatalf("unable to open database connection: %b", err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatalf("unable to connect to database %b", err)
	}
	h := &api.Handler{DB: db}
	defer db.Close()

	t, err := template.ParseGlob("templates/*/*.html")
	if err != nil {
		e.Logger.Fatalf("unable to parse templates: %b", err)
	}
	e.Renderer = &Templates{templates: t}

	e.Static("/css", "public/css")
	e.Static("/js", "public/js")
	e.Static("/static", "public/assets")

	// pages
	e.GET("/", h.Index)
	e.GET("/add", h.Add)

	// recipe
	e.GET("/recipe/:id", h.GetRecipe)
	e.POST("/recipe", h.PostRecipe)
	// list
	e.GET("/list/:id", h.GetList)
	e.POST("/list_search", h.GetLists)
	// list item
	e.GET("/item/:id", h.GetItem)
	e.PUT("/item/:type/:id", h.PutItem)
	e.GET("/item/:id/edit", h.EditItem)
	e.DELETE("/item/:type/:id", h.DeleteItem)

	e.Debug = true
	e.Logger.Fatal(e.Start(":1323"))
}
