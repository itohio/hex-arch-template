package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	ErrExit = errors.New("Simple exit")
)

type Config struct {
	Context context.Context
	Logger  *log.Logger

	HostName     string
	FrontEndPath string
	Server       Server
	Auth         Auth
	Db           Db
	NATS         NATS
}

type Server struct {
	Origins []string
	Host    string
	Port    int
}

type Auth struct {
	Domain   string
	ClientID string
	AuthKey  string
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
	viper.SetDefault("ORIGINS", []string{"localhost", "127.0.0.1"})
	viper.SetDefault("HOST", "localhost")
	viper.SetDefault("PORT", 8080)

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
		Server: Server{
			Origins: v.GetStringSlice("ORIGINS"),
			Host:    v.GetString("HOST"),
			Port:    v.GetInt("PORT"),
		},
		Auth: Auth{
			Domain:   v.GetString("AUTH0_DOMAIN"),
			ClientID: v.GetString("AUTH0_CLIENT_ID"),
			AuthKey:  v.GetString("AUTH0_KEY"),
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

	return nil
}
