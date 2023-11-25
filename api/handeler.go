package api

import (
	"database/sql"
	"fmt"
	"htmx-test/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Handler struct {
		DB *sql.DB
	}
)

func (handler Handler) Index(c echo.Context) error {
	lists, err := db.GetTopLevelLists(handler.DB)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%b", err))
	}
	return c.Render(http.StatusOK, "home", lists)
}

func (handler Handler) Add(c echo.Context) error {
	return c.Render(http.StatusOK, "add", nil)
}
