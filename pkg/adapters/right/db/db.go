package pgql

import (
	"fmt"
	"hexarch/pkg/config"
	"hexarch/pkg/models"
	"hexarch/pkg/ports"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Adapter struct {
	cfg *config.Config
	db  *gorm.DB
}

var _ ports.DbPort = &Adapter{}

func New(cfg *config.Config, migrate bool) (*Adapter, error) {
	var driver gorm.Dialector
	if cfg.Db.Host == "SQLite" {
		driver = sqlite.Open(cfg.Db.Name)
	} else {
		conn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.Db.Host,
			cfg.Db.Port,
			cfg.Db.Username,
			cfg.Db.Password,
			cfg.Db.Name,
			cfg.Db.SSLMode,
		)
		driver = postgres.Open(conn)
	}

	db, err := gorm.Open(driver, &gorm.Config{})

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	if migrate {
		db.AutoMigrate(
			&models.Greeting{},
		)
	}

	return &Adapter{
		cfg: cfg,
		db:  db,
	}, nil
}

func (a *Adapter) GetRandomGreeting() string {
	return ""
}

func (a *Adapter) GetGreetings() []string {
	return nil
}
