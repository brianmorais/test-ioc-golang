package repository_interfaces

import "github.com/moraisbrian/reservations-sale/domain/entities"

type StationRepositoryInterface interface {
	Save(s entities.Station) error
}
