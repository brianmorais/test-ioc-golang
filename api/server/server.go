package server

import (
	"github.com/labstack/echo/v4"
	"github.com/brianmorais/reservations-sale/api/server/handlers"
)

func Serve() {
	e := echo.New()
	e.POST("/stations", handlers.ServiceHandler)
	e.Logger.Fatal(e.Start(":3000"))
}
