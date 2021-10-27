package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthcheckResult struct {
	// pass / fail / warn
	Status string
}

func Healthcheck(e echo.Context) error {
	status := HealthcheckResult{
		Status: "pass",
	}

	return e.JSON(http.StatusOK, status)
}
