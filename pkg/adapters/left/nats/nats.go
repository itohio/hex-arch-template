package nats

import (
	"fmt"
	"hexarch/pkg/config"
	"hexarch/pkg/ports"

	"github.com/nats-io/nats.go"
)

type Adapter struct {
	cfg *config.Config
	nc  *nats.Conn
	app ports.APIPort
}

func New(cfg *config.Config, app ports.APIPort) (*Adapter, error) {
	// Connect Options.
	opts := []nats.Option{
		nats.Name(cfg.HostName),
		nats.UserInfo(cfg.NATS.Username, cfg.NATS.Password),
	}

	// Connect to NATS
	nc, err := nats.Connect(cfg.NATS.URL, opts...)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		cfg: cfg,
		nc:  nc,
		app: app,
	}, nil
}

func (a *Adapter) Run() {
	topic := fmt.Sprintf("%s.hello", a.cfg.NATS.TopicID)
	a.nc.Subscribe(topic, func(msg *nats.Msg) {
		a.nc.Publish(topic, []byte(a.app.SayHello(string(msg.Data))))
	})

	return
}
