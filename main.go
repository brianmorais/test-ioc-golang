package main

import (
	"github.com/moraisbrian/reservations-sale/api/server"
	"github.com/moraisbrian/reservations-sale/ioc"
)

func main() {
	ioc.SetDependencies()
	server.Serve()
}
