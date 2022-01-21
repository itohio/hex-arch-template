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

	Server Server
	Auth   Auth
	Db     Db
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
		Db: Db{
			Username: v.GetString("DB_USERNAME"),
			Password: v.GetString("DB_PASSWORD"),
			Url:      v.GetString("DB_URL"),
		},
	}

	return nil
}
