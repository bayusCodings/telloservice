package telloservice

import (
	"fmt"
	"os"
	"time"

	"github.com/bayuscodings/telloservice/app/auth"
	"github.com/bayuscodings/telloservice/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ApplicationHandler struct {
	Config *config.Configuration
	DB     *gorm.DB
	JWT    auth.TokenMaker
}

func (a *ApplicationHandler) PrepareConfig() {
	a.Config = config.New()
}

func (a *ApplicationHandler) InitDB() {
	config := a.Config.Database
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.DbHost, config.DbUsername, config.DbPassword, config.DbName, config.DbPort)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		defer fmt.Printf("DATABASE Error: %v", err)
	}

	a.DB = db

	if config.EnableLog {
		a.DB = db.Debug()
	}
}

func (a *ApplicationHandler) InitTokenMaker() {
	tokenMaker := auth.NewJWTMaker(a.Config.SecretKey)
	a.JWT = tokenMaker
}

func (a *ApplicationHandler) InitLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
}
