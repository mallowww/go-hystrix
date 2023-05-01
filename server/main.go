package main

import (
	"net/http"

	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
	return c.String(http.StatusOK, "home")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", home)
	e.Logger.Fatal(e.Start(":1322"))
}
