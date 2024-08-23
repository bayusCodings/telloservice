package config

import (
	"os"
	"strconv"
)

type DatabaseConfig struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUsername string
	DbPassword string
	EnableLog  bool
	RunSeeding bool
}

type Configuration struct {
	AppPort   string
	AppEnv    string
	SecretKey string
	Database  *DatabaseConfig
}

func getEnvValue(orginal string, alternative string) string {
	if orginal == "" {
		return alternative
	}

	return orginal
}

func getEnvBoolValue(orginal string, alternative string) bool {
	value := orginal
	if value == "" {
		value = alternative
	}

	if b, err := strconv.ParseBool(value); err != nil {
		return false
	} else {
		return b
	}
}

func New() *Configuration {
	c := new(Configuration)

	c.AppPort = getEnvValue(os.Getenv("APP_PORT"), "50055")
	c.AppEnv = getEnvValue(os.Getenv("APP_ENV"), "local")
	c.SecretKey = getEnvValue(os.Getenv("SECRET_KEY"), "supersecret")

	c.Database = new(DatabaseConfig)
	c.Database.DbHost = getEnvValue(os.Getenv("DB_HOST"), "127.0.0.1")
	c.Database.DbPort = getEnvValue(os.Getenv("DB_PORT"), "3306")
	c.Database.DbName = getEnvValue(os.Getenv("DB_DATABASE"), "telloservice")
	c.Database.DbUsername = getEnvValue(os.Getenv("DB_USERNAME"), "root")
	c.Database.DbPassword = getEnvValue(os.Getenv("DB_PASSWORD"), "secret")
	c.Database.EnableLog = getEnvBoolValue(os.Getenv("DB_LOGGING"), "0")
	c.Database.RunSeeding = getEnvBoolValue(os.Getenv("RUN_SEEDER"), "1")

	return c
}
