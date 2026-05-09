package main

import (
	"fmt"
	"gnet/internal/config"
	"gnet/internal/tools"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	if err := config.Init(); err != nil {
		log.Fatalf("Ошибка при загрузке .env файла")
	}
	envData := config.GetEnvData()
	models, err := tools.SendGetWithContext(envData.APIUrl, envData.APIKey, client)
	if err != nil {
		log.Fatalf("Получен отрицательный ответ с кодом не 200")
	}
	fmt.Print(models)
}
