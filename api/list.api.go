package api

import (
	"fmt"
	"gourmeg/db"
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
	if id <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id: (%d)", id))
	}

	l, err := db.GetList(handler.DB, int(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%b", err))
	}

	return c.Render(http.StatusOK, "list.html", l)
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

func (handler Handler) PostList(c echo.Context) error {
	id_str := c.QueryParam("parent_id")

	parent_id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("parent_id param (%s) not a number", id_str))
	}

	name := c.FormValue("name")
	if len(name) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "list name cannot be empty")
	}

	if len(name) > 512 {
		return echo.NewHTTPError(http.StatusBadRequest, "list name cannot be longer than 512 characters")
	}

	item, err := db.PostList(handler.DB, name, parent_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}

	item.ListIds = []int{parent_id}

	return c.Render(http.StatusOK, "add_item_options_and_items", db.List{Id: parent_id, Children: []db.Item{item}})
}
