package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func home(counter int) func(c echo.Context) error {
	return func(c echo.Context) error {
		counter++

		if counter >= 5 && counter <= 10 {
			time.Sleep(time.Millisecond * 1000)
		}
		msg := fmt.Sprintf("Hello %v", counter)
		fmt.Println(msg)

		return c.String(http.StatusOK, msg)
	}
}

func main() {
	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	counter := 0
	e.GET("/api", home(counter))
	e.Logger.Fatal(e.Start(":8000"))
}
