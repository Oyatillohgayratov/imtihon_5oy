package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Postgres PostgresConfig

	ApigatewayPort string

	UserServiceHost string
	UserServicePort string

	HotelServiceHost string
	HotelServicePort string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Load(path string) Config {
	err := godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatalf("Error loading.env file: %s", err)
	}
	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		Postgres: PostgresConfig{
			Host:     conf.GetString("POSTGRES_HOST"),
			Port:     conf.GetString("POSTGRES_PORT"),
			User:     conf.GetString("POSTGRES_USER"),
			Password: conf.GetString("POSTGRES_PASSWORD"),
			Database: conf.GetString("POSTGRES_DATABASE"),
		},

		ApigatewayPort: conf.GetString("API_GATEWAY_PORT"),

		UserServiceHost: conf.GetString("USER_SERVICE_HOST"),
		UserServicePort: conf.GetString("USER_SERVICE_PORT"),

		HotelServiceHost: conf.GetString("HOTEL_SERVICE_HOST"),
		HotelServicePort: conf.GetString("HOTEL_SERVICE_PORT"),
	}

	return cfg
}
