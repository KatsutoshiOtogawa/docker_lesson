package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		jsonMap := map[string]string{
			"foo":  "bar",
			"hoge": "fuga",
		}

		return c.JSON(http.StatusOK, jsonMap)

		// e.Use(standard.WrapMiddleware(cors.New(cors.Options{
		// 	AllowedOrigins: []string{"http://example.com"},
		// }).Handler))

	})
	e.Logger.Fatal(e.Start(":8082"))
}
