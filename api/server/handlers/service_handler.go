package handlers

import (
	"net/http"

	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
	"github.com/brianmorais/reservations-sale/api/controllers"
	"github.com/brianmorais/reservations-sale/application/dtos"
	service_interfaces "github.com/brianmorais/reservations-sale/application/interfaces"
)

func ServiceHandler(c echo.Context) error {
	dto := new(dtos.StationDto)

	if err := c.Bind(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var err error
	var stationService service_interfaces.StationServiceInterface
	if err = container.Resolve(&stationService); err != nil {
		panic("error")
	}

	controller := controllers.NewStationController(stationService)

	if err = controller.SaveStation(*dto); err == nil {
		return c.JSON(http.StatusCreated, dto)
	}

	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}
