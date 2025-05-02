package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	RestPort     string      `env:"REST_PORT" envDefault:"8080"`
	AllowOrigins []string    `env:"ALLOW_ORIGINS" envSeparator:","`
	LogFormat    string      `env:"LOG_FORMAT"`
	Mongo        MongoConfig `envPrefix:"MONGO_"`
}

type MongoConfig struct {
	URI      string `env:"URI"`
	Database string `env:"DATABASE"`
}

// @WireSet("Config")
func NewConfig() *Config {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it. Falling back to system environment variables.")
	}

	config := &Config{}

	// Parse environment variables into the config struct
	if err := env.Parse(config); err != nil {
		log.Fatalln("Failed to parse environment variables into Config struct:", err)
	}

	return config
}
