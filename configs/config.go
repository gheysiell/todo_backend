package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	viper.AutomaticEnv()
}

func Load() error {
	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("API_PORT"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Pass:     viper.GetString("DB_PASS"),
		Database: viper.GetString("DB_NAME"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
