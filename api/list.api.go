package api

import (
	"fmt"
	"htmx-test/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (handler Handler) GetList(c echo.Context) error {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
	}

	l, err := db.GetList(handler.DB, int(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%b", err))
	}

	return c.Render(http.StatusOK, "list_page", l)
}

func (handler Handler) GetLists(c echo.Context) error {
	list := c.FormValue("list")

	items, err := db.SearchList(handler.DB, list)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}

	return c.Render(http.StatusOK, "list_search_results", items)
}

func (handler Handler) EditList(c echo.Context) error {
    id_str := c.Param("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
	}

	l, err := db.GetList(handler.DB, int(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("%b", err))
	}
    c.Logger().Infof("%s", l)

	return c.Render(http.StatusOK, "edit_list", l)
}

func (handler Handler) DeleteList(c echo.Context) error {
	id_str := c.Param("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
	}

	err = db.DeleteList(handler.DB, int(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}

    return c.Redirect(http.StatusSeeOther, "/")
}
