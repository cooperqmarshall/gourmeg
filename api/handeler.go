package api

import (
	"database/sql"
	"fmt"
	"gourmeg-htmx/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Handler struct {
		DB *sql.DB
	}
)

func (handler Handler) Index(c echo.Context) error {
    root_list, err := db.GetList(handler.DB, 1) // TODO: get root list id per user
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%b", err))
	}
	return c.Render(http.StatusOK, "index.html", root_list.Children)
}

func (handler Handler) Add(c echo.Context) error {
	return c.Render(http.StatusOK, "add", nil)
}
