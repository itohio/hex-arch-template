package db

import (
	"hexarch/pkg/config"
	"hexarch/pkg/ports"
)

type Adapter struct {
	cfg *config.Config
}

var _ ports.DbPort = &Adapter{}

func New(cfg *config.Config) *Adapter {
	return &Adapter{
		cfg: cfg,
	}
}

func (a *Adapter) GetRandomGreeting() string {
	return "Hello,"
}

func (a *Adapter) GetGreetings() []string {
	return []string{"Hello,", "Hi,", "Ahoy,"}
}
