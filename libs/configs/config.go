package configs

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Host          string `env:"APP_HOST"`
	Port          int    `env:"APP_PORT"`
}

type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST"`
	Port     int    `env:"DATABASE_PORT"`
	Username string `env:"DATABASE_USERNAME"`
	Password string `env:"DATABASE_PASSWORD"`
	Name     string `env:"DATABASE_NAME"`
}

type config struct {
	App   AppConfig
	DB    DatabaseConfig
}

var (
	App   AppConfig
	DB    DatabaseConfig
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file: " + err.Error())
	}

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Error parse env data to struct: " + err.Error())
	}

	App = cfg.App
	DB = cfg.DB
}
