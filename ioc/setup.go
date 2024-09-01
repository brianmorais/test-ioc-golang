package ioc

import (
	"github.com/golobby/container/v3"
	service_interfaces "github.com/brianmorais/reservations-sale/application/interfaces"
	"github.com/brianmorais/reservations-sale/application/services"
	"github.com/brianmorais/reservations-sale/data/repositories"
	repository_interfaces "github.com/brianmorais/reservations-sale/domain/interfaces"
)

func SetDependencies() {
	container.Transient(func() repository_interfaces.StationRepositoryInterface {
		return &repositories.StationRepository{}
	})

	container.Transient(func() service_interfaces.StationServiceInterface {
		return &services.StationService{}
	})
}
