package service_interfaces

import "github.com/brianmorais/reservations-sale/application/dtos"

type StationServiceInterface interface {
	Save(s dtos.StationDto) error
}
