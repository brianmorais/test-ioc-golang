package main

import (
	"github.com/brianmorais/reservations-sale/api/server"
	"github.com/brianmorais/reservations-sale/ioc"
)

func main() {
	ioc.SetDependencies()
	server.Serve()
}
