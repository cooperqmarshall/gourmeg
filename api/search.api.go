package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (handler Handler) Search(c echo.Context) error {
    return c.Render(http.StatusOK, "search", nil)
}
