package helloworld

import (
	"fmt"
	"hexarch/pkg/app/api"
	"hexarch/pkg/config"
)

type HelloWorld struct {
	cfg *config.Config
}

// Check if we actually implement relevant api
var _ api.HelloWorld = &HelloWorld{}

func New(cfg *config.Config) *HelloWorld {
	return &HelloWorld{
		cfg: cfg,
	}
}

func (h *HelloWorld) HelloWorld(greeting, name string) string {
	return fmt.Sprintf("%s %s", greeting, name)
}
