package main

import (
	"fmt"
	"os"
	"sample-pod-image/handler"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := 80
	if os.Getenv("PORT") != "" {
		portVal, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			panic(err)
		}
		port = portVal
	}

	run(port)
}

func run(port int) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", handler.Index)
	e.GET("/health", handler.Healthcheck)

	bind := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(bind))
}
