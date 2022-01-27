package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	EnvPrefix = "HEX"
	Version   = "0.1.0"
	Build     = ""
)

var (
	ErrExit = errors.New("Simple exit")
)

type Config struct {
	Context context.Context
	Logger  *log.Logger

	HostName     string
	FrontEndPath string
	Debug        bool
	Server       Server
	GRPC         GRPC
	Auth         Auth
	Db           Db
	NATS         NATS
}

type Server struct {
	Origins []string
	Host    string
	Address string
}

type GRPC struct {
	Network string
	Address string
}

type Auth struct {
	Domain   string
	Audience string
	ClientID string
	Secret   string
}

type NATS struct {
	URL      string
	Username string
	Password string
	TopicID  string
}

type Db struct {
	Username string
	Password string
	Host     string
	Name     string
	Port     int
	SSLMode  string
}

func isDotEnvPresent() bool {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return false
	}
	return true
}

func New(ctx context.Context) *Config {
	// viper.SetEnvPrefix(EnvPrefix)

	viper.SetDefault("ORIGINS", "http://localhost:8080,http://localhost:3000")
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("ADDRESS", ":8080")
	viper.SetDefault("GRPC_NETWORK", "tcp")
	viper.SetDefault("GRPC_ADDRESS", ":9000")

	return &Config{
		Context: ctx,
		Logger:  log.Default(),
	}
}

func (c *Config) Read() error {
	if isDotEnvPresent() {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env")
		viper.SetConfigType("dotenv")

		err := viper.ReadInConfig()
		if err != nil {
			return err
		}
	}
	viper.AutomaticEnv()

	return c.parseConfig(viper.GetViper())
}

func (c *Config) Write() error {
	return fmt.Errorf("not implemented")
}

func (c *Config) parseConfig(v *viper.Viper) error {
	*c = Config{
		HostName:     v.GetString("HOSTNAME"),
		FrontEndPath: v.GetString("FRONT_END_PATH"),
		Debug:        v.GetBool("DEBUG"),
		Server: Server{
			Origins: strings.Split(v.GetString("ORIGINS"), ","),
			Host:    v.GetString("HOST"),
			Address: v.GetString("ADDRESS"),
		},
		GRPC: GRPC{
			Network: v.GetString("GRPC_NETWORK"),
			Address: v.GetString("GRPC_ADDRESS"),
		},
		Auth: Auth{
			Domain:   v.GetString("AUTH0_DOMAIN"),
			Audience: v.GetString("AUTH0_AUDIENCE"),
			ClientID: v.GetString("AUTH0_CLIENT_ID"),
			Secret:   v.GetString("AUTH0_SECRET"),
		},
		NATS: NATS{
			URL:      v.GetString("NATS_URL"),
			Username: v.GetString("NATS_USERNAME"),
			Password: v.GetString("NATS_PASSWORD"),
			TopicID:  v.GetString("NATS_TOPIC_ID"),
		},
		Db: Db{
			Username: v.GetString("DB_USERNAME"),
			Password: v.GetString("DB_PASSWORD"),
			Host:     v.GetString("DB_HOST"),
			Port:     v.GetInt("DB_PORT"),
			Name:     v.GetString("DB_NAME"),
			SSLMode:  v.GetString("DB_SSL_MODE"),
		},
	}

	if c.Debug {
		log.Println("WARNING: DEBUG Mode enabled!")
	}

	return nil
}
