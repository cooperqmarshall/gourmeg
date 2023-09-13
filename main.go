package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"htmx-test/api"
)

type Templates struct {
  templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}
 
func main() { 
  t, err := template.ParseGlob("public/views/*.html")

  if err != nil {
    log.Fatalf("unable to parse templates: %b", err)
  }

  e := echo.New()
  e.Renderer = &Templates{templates: t}

  e.Use(middleware.Logger())
  e.Static("/api", "public/api")
  e.Static("/css", "public/css")

  e.GET("/", func(c echo.Context) error {
    return c.Render(http.StatusOK, "index.html", nil)
  })

  e.GET("/book/:id", func(c echo.Context) error {
    id_str := c.Param("id")
    id, err := strconv.ParseInt(id_str, 10, 0)
    if err != nil {log.Fatalf("id param not a number: %b", err)}

    api.GetBook(int(id))
    return c.Render(http.StatusOK, "index.html", nil)
  })

  e.GET("/chapter/:id", func(c echo.Context) error {
    id_str := c.Param("id")
    id, err := strconv.ParseInt(id_str, 10, 0)
    if err != nil {log.Fatalf("id param not a number: %b", err)}

    api.GetChapter(int(id))
    return c.Render(http.StatusOK, "index.html", nil)
  })

  e.GET("/recipe/:id", func(c echo.Context) error {
    id_str := c.Param("id")
    id, err := strconv.ParseInt(id_str, 10, 0)
    if err != nil {log.Fatalf("id param not a number: %b", err)}

    recipe := api.GetRecipe(int(id))
    _ = recipe
    return c.Render(http.StatusOK, "recipe_page", recipe)
  })

  e.Logger.Fatal(e.Start(":1323"))
}

