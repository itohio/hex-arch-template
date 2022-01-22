package main

import (
	"context"
	"errors"
	"fmt"
	"hexarch/pkg/adapters/left/auth"
	"hexarch/pkg/adapters/left/gql"
	"hexarch/pkg/adapters/left/grpc"
	"hexarch/pkg/adapters/left/nats"
	"hexarch/pkg/adapters/left/spa"
	"hexarch/pkg/adapters/right/memory"
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
	db := memory.New(cfg)

	// Create the Domain layer and Application layer
	hw := helloworld.New(cfg)
	app := api.New(cfg, db, hw)

	// Create left adapters
	authRouter := auth.New(cfg)
	singlePageApp := spa.New(cfg.FrontEndPath)
	gql, err := gql.New(cfg, app, db, map[string]http.Handler{
		"/auth": authRouter,
		"/":     singlePageApp,
	})
	if err != nil {
		panic(err)
	}
	go gql.Run()

	rpc, err := grpc.New(cfg, app)
	if err != nil {
		fmt.Println("Could not create GRPC: ", err)
	} else {
		go rpc.Run()
	}

	nc, err := nats.New(cfg, app)
	if err != nil {
		fmt.Println("Could not create NATS: ", err)
	} else {
		go nc.Run()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-c
}
