package api

import (
	"database/sql"
	"fmt"
	"gourmeg/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Handler struct {
		DB *sql.DB
	}
)

func (handler Handler) Index(c echo.Context) error {
	l := db.ListTree{Id:0}
	err := db.GetListTree(handler.DB, &l) // TODO: get root list id per user
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%b", err))
	}
	return c.Render(http.StatusOK, "index.html", l)
}

func (handler Handler) Add(c echo.Context) error {
	return c.Render(http.StatusOK, "add.html", nil)
}
