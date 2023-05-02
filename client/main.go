package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/api", api)
	e.Logger.Fatal(e.Start(":8001"))
}

func init() {
	hystrix.ConfigureCommand("api", hystrix.CommandConfig{
		Timeout:                500,
		RequestVolumeThreshold: 1,
		ErrorPercentThreshold:  100,
		SleepWindow:            15000,
	})

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(":8002", hystrixStreamHandler)
}

func api(c echo.Context) error {
	hystrix.Go("api", func() error {

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

		return nil
	}, func(err error) error {
		fmt.Println(err)
		return nil
	})

	return nil
}
