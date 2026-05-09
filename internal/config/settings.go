package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvData struct {
	APIKey string `env:"APIKEY"`
	APIUrl string `env:"APIURL"`
}

func Init() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка при загрузке .env файла: %s", err)
	}
	return nil
}
func GetEnvData() EnvData {
	env_data := EnvData{os.Getenv("APIKEY"), os.Getenv("APIURL")}
	return env_data
}
