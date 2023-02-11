package ioc

import (
	"github.com/golobby/container/v3"
	service_interfaces "github.com/moraisbrian/reservations-sale/application/interfaces"
	"github.com/moraisbrian/reservations-sale/application/services"
	"github.com/moraisbrian/reservations-sale/data/repositories"
	repository_interfaces "github.com/moraisbrian/reservations-sale/domain/interfaces"
)

func SetDependencies() {
	container.Transient(func() repository_interfaces.StationRepositoryInterface {
		return &repositories.StationRepository{}
	})

	container.Transient(func() service_interfaces.StationServiceInterface {
		return &services.StationService{}
	})
}
