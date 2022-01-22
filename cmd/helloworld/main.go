package main

import (
	"context"
	"errors"
	"hexarch/pkg/adapters/left/auth"
	"hexarch/pkg/adapters/left/gql"
	"hexarch/pkg/adapters/left/spa"
	"hexarch/pkg/adapters/right/db"
	"hexarch/pkg/app/api"
	"hexarch/pkg/app/core/helloworld"
	"hexarch/pkg/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := config.New(ctx)

	if err := cfg.Read(); err != nil {
		if errors.Is(err, config.ErrExit) {
			return
		}
		panic(err)
	}

	// Create Right adapters
	db := db.New(cfg)

	// Create the Domain layer and Application layer
	hw := helloworld.New(cfg)
	app := api.New(cfg, db, hw)

	// Create left adapters
	authRouter := auth.New(cfg)
	singlePageApp := spa.New(cfg.FrontEndPath)
	gql := gql.New(cfg, app, map[string]http.Handler{
		"auth/": authRouter,
		"/":     singlePageApp,
	})
	go gql.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-c
}
