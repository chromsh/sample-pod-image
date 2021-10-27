package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]string{"hello": "world"})
}
