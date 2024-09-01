package repository_interfaces

import "github.com/brianmorais/reservations-sale/domain/entities"

type StationRepositoryInterface interface {
	Save(s entities.Station) error
}
