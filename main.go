package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// health check struct
type HealthCheckResponse struct {
	Status string `json:"status"`
}

func main() {
	// new Echo server instance
	e := echo.New()

	// root route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Echo!")
	})

	// health check endpoint
	e.GET("/healthy", func(c echo.Context) error {
		h := HealthCheckResponse{Status: "OK"}
		return c.JSON(http.StatusOK, h)
	})

	// ping path
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, HealthCheckResponse{Status: "PONG!!"})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
