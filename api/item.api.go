package api

import (
	"fmt"
	"gourmeg/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (handler Handler) GetItem(c echo.Context) error {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
	}
	var i db.Item
	c.Bind(&i)

	i, err = db.GetItem(handler.DB, int(id), i.Type)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}
	return c.Render(http.StatusOK, "item", i)
}

func (handler Handler) PutItem(c echo.Context) error {
	var i db.Item
	err := c.Bind(&i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}

	recipe, err := db.UpdateItem(handler.DB, i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}
	return c.Render(http.StatusOK, "item", recipe)
}

func (handler Handler) EditItem(c echo.Context) error {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
	}

	var i db.Item
	err = c.Bind(&i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
	}
	i.Id = int(id)

	return c.Render(http.StatusOK, "edit_item", i)
}

func (handler Handler) DeleteItem(c echo.Context) error {
	var i db.Item
	err := c.Bind(&i)
	fmt.Printf("%s\n", i.Type)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}

	err = db.DeleteItem(handler.DB, i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}
	return c.NoContent(http.StatusOK)
}

func (handler Handler) AddListItem(c echo.Context) error {
	id_str := c.QueryParam("parent_id")

	parent_id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("parent_id param (%s) not a number", id_str))
	}

	return c.Render(http.StatusOK, "add_list_item", parent_id)
}

func (handler Handler) AddRecipeItem(c echo.Context) error {
	id_str := c.QueryParam("list_id")

	list_id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("list_id param (%s) not a number", id_str))
	}

	return c.Render(http.StatusOK, "add_recipe_item", list_id)
}

func (handler Handler) ItemSearch(c echo.Context) error {
	search_term := c.FormValue("search_term")
	search_results, err := db.ItemSearch(handler.DB, search_term)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.Render(http.StatusOK, "list_of_items", search_results)
}

func (handler Handler) AddItemOptions(c echo.Context) error {
	id_str := c.QueryParam("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("list_id param (%s) not a number", id_str))
	}
	
	return c.Render(http.StatusOK, "add_item_options", id)
}
