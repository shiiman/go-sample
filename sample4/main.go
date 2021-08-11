package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"Code": http.StatusOK, "Message": "health"})
	})
	e.Logger.Fatal(e.Start(":8080"))
}