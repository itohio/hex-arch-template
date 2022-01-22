package grpc

//go:generate protoc -I=proto -I=${GOPATH}/pkg/mod/ -I=${GOPATH}/src --gogofaster_out=plugins=grpc:. schema.proto

import (
	"context"
	"fmt"
	"log"
	"net"

	"hexarch/pkg/adapters/left/grpc/pb"
	"hexarch/pkg/config"
	"hexarch/pkg/ports"

	"google.golang.org/grpc"
)

type Adapter struct {
	cfg    *config.Config
	api    ports.APIPort
	listen net.Listener
}

func New(cfg *config.Config, api ports.APIPort) (*Adapter, error) {
	ret := &Adapter{api: api}

	listen, err := net.Listen(cfg.GRPC.Network, cfg.GRPC.Address)
	if err != nil {
		return nil, err
	}
	ret.listen = listen

	return ret, nil
}

func (a *Adapter) Run() {
	grpcServer := grpc.NewServer()
	pb.RegisterHelloWorldServer(grpcServer, a)
	if err := grpcServer.Serve(a.listen); err != nil {
		log.Fatalf("failed to serve gRPC server over address %s: %v", a.cfg.GRPC.Address, err)
	}
}

func (a *Adapter) GetGreeting(ctx context.Context, input *pb.Input) (*pb.Answer, error) {
	if input == nil {
		return nil, fmt.Errorf("input is mandatory")
	}

	return &pb.Answer{
		Greeting: a.api.SayHello(input.Name),
	}, nil
}
