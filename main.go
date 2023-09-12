package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

  e.Logger.Fatal(e.Start(":1323"))
}

