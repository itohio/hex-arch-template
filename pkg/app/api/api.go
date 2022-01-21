package api

import (
	"hexarch/pkg/config"
	"hexarch/pkg/ports"
)

type Application struct {
	cfg     *config.Config
	db      ports.DbPort
	greeter HelloWorld
}

// Check if we actually implement all the ports.
var _ ports.APIPort = &Application{}

func New(cfg *config.Config, db ports.DbPort, greeter HelloWorld) *Application {
	return &Application{
		cfg:     cfg,
		greeter: greeter,
		db:      db,
	}
}

func (a *Application) SayHello(name string) string {
	return a.greeter.HelloWorld(a.db.GetRandomGreeting(), name) + "!"
}
