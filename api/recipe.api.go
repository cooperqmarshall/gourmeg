package api

import (
	"fmt"
	"htmx-test/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (handler Handler) PostRecipe(c echo.Context) error {
	r := new(db.Recipe)
	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := db.PostRecipe(handler.DB, r)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}

	return c.Render(http.StatusOK, "add_recipe_result", r)
}

func (handler Handler) GetRecipe(c echo.Context) error {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
	}

	recipe, err := db.GetRecipe(handler.DB, int(id))
	fmt.Printf("%d, %d", id, recipe.Id)
	return c.Render(http.StatusOK, "recipe_page", recipe)
}
