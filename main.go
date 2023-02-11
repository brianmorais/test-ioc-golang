package main

import (
	"github.com/golobby/container/v3"
	"github.com/moraisbrian/reservations-sale/api/controllers"
	"github.com/moraisbrian/reservations-sale/application/dtos"
	service_interfaces "github.com/moraisbrian/reservations-sale/application/interfaces"
	"github.com/moraisbrian/reservations-sale/ioc"
)

func main() {
	ioc.SetDependencies()

	var err error
	var stationService service_interfaces.StationServiceInterface
	if err = container.Resolve(&stationService); err != nil {
		panic("error")
	}

	controller := controllers.NewStationController(stationService)

	dto := dtos.StationDto{
		Id:   10,
		Name: "REC",
	}

	controller.SaveStation(dto)
}
