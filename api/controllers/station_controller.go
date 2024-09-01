package controllers

import (
	"github.com/golobby/container/v3"
	"github.com/brianmorais/reservations-sale/application/dtos"
	service_interfaces "github.com/brianmorais/reservations-sale/application/interfaces"
	"github.com/brianmorais/reservations-sale/application/services"
	repository_interfaces "github.com/brianmorais/reservations-sale/domain/interfaces"
)

type StationController struct {
	StationService service_interfaces.StationServiceInterface
}

func NewStationController(stationService service_interfaces.StationServiceInterface) *StationController {
	return &StationController{
		StationService: stationService,
	}
}

func (c *StationController) SaveStation(stationDto dtos.StationDto) error {
	var err error
	var stationRepository repository_interfaces.StationRepositoryInterface
	if err = container.Resolve(&stationRepository); err != nil {
		panic("error")
	}

	service := services.NewStationService(stationRepository)
	err = service.Save(stationDto)
	return err
}
