package repositories

import (
	"fmt"

	"github.com/moraisbrian/reservations-sale/domain/entities"
)

type StationRepository struct{}

func (s *StationRepository) Save(station entities.Station) error {
	fmt.Printf("chegou no repositorio, valor da station: %v - %s", station.Id, station.Name)
	return nil
}
