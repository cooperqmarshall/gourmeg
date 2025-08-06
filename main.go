package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"

	"gourmeg/api"
)

type Templates struct {
	templates []*template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    fmt.Println(name)
	switch name {
    case "index.html":
        return t.templates[1].ExecuteTemplate(w, "base", data)
    case "list.html":
        return t.templates[2].ExecuteTemplate(w, "base", data)
    case "recipe.html":
        return t.templates[3].ExecuteTemplate(w, "base", data)
    case "add.html":
        return t.templates[4].ExecuteTemplate(w, "base", data)
    case "search.html":
        return t.templates[5].ExecuteTemplate(w, "base", data)
	default:
		return t.templates[0].ExecuteTemplate(w, name, data)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	postgres_uri := os.Getenv("POSTGRESQL_URI")
	if len(postgres_uri) == 0 {
		postgres_uri = "host=localhost user=postgres password=secret dbname=gourmegdb sslmode=disable"
	}

	db, err := sql.Open("postgres", postgres_uri)
	if err != nil {
		e.Logger.Fatalf("unable to open database connection: %b", err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatalf("unable to connect to database %b", err)
	}
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(1000 * time.Millisecond)
	h := &api.Handler{DB: db}
	defer db.Close()

	templates_dir := os.Getenv("TEMPLATES_DIR")
	if len(templates_dir) == 0 {
		templates_dir = "templates/*.html"
	}

    var templates []*template.Template
	t := template.Must(template.ParseGlob(templates_dir))
    templates = append(templates, t)

    t = template.Must(template.Must(t.Clone()).ParseFiles("templates/index.html", "templates/_base.html"))
    templates = append(templates, t)

    t = template.Must(template.Must(t.Clone()).ParseFiles("templates/list.html", "templates/_base.html"))
    templates = append(templates, t)

    t = template.Must(template.Must(t.Clone()).ParseFiles("templates/recipe.html", "templates/_base.html"))
    templates = append(templates, t)

    t = template.Must(template.Must(t.Clone()).ParseFiles("templates/add.html", "templates/_base.html"))
    templates = append(templates, t)

    t = template.Must(template.Must(t.Clone()).ParseFiles("templates/search.html", "templates/_base.html"))
    templates = append(templates, t)

	e.Renderer = &Templates{templates: templates}

	e.Static("/css", "public/css")
	e.Static("/js", "public/js")
	e.Static("/static", "public/assets")

	// pages
	e.GET("/", h.Index)
	e.GET("/add", h.Add)
	e.GET("/search", h.Search)

	// recipe
	e.GET("/recipe/:id", h.GetRecipe)
	e.POST("/recipe", h.PostRecipe)
	e.PUT("/recipe/refetch/:id", h.RefetchRecipe)

	// list
	e.GET("/list/:id", h.GetList)
	e.GET("/list/:id/edit", h.EditList)
	e.DELETE("/list/:id", h.DeleteList)
	e.POST("/list", h.PostList)
	e.POST("/list_search", h.GetLists)

	// list item
	e.GET("/item/:id", h.GetItem)
	e.PUT("/item/:type/:id", h.PutItem)
	e.GET("/item/:id/edit", h.EditItem)
	e.GET("/item/recipe/add", h.AddRecipeItem)
	e.GET("/item/list/add", h.AddListItem)
	e.DELETE("/item/:type/:id", h.DeleteItem)
	e.POST("/item/search", h.ItemSearch)

	e.Debug = true
	e.Logger.Fatal(e.Start(":1323"))
}
