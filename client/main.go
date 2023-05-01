package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	
	e.GET("/api", api)
	e.Logger.Fatal(e.Start(":8001"))
}

func api(c echo.Context) error {
	res, err := http.Get("http://localhost:8000/api")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	msg := string(data)
	fmt.Println(msg)

	return c.String(http.StatusOK, msg)
}
