package graph

import (
	"hexarch/pkg/config"
	"hexarch/pkg/ports"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Cfg *config.Config
	App ports.APIPort
	Db  ports.DbPort
}
