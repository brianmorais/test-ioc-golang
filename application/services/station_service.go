package services

import (
	"github.com/brianmorais/reservations-sale/application/dtos"
	"github.com/brianmorais/reservations-sale/domain/entities"
	repository_interfaces "github.com/brianmorais/reservations-sale/domain/interfaces"
)

type StationService struct {
	StationRepository repository_interfaces.StationRepositoryInterface
}

func NewStationService(stationRepository repository_interfaces.StationRepositoryInterface) StationService {
	return StationService{
		StationRepository: stationRepository,
	}
}

func (s *StationService) Save(stationDto dtos.StationDto) error {
	station := entities.Station{
		Id:   stationDto.Id,
		Name: stationDto.Name,
	}

	err := s.StationRepository.Save(station)
	return err
}
