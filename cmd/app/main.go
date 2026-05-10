package main

import (
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
	data := make(map[string]any)
	data["model"] = "gpt-3.5-turbo"
	data["messages"] = []any{
		map[string]any{
			"role":    "user",
			"content": "Привет!",
		},
	}
	data["stream"] = false

	tools.SendPostWithContext(envData.APIUrl, envData.APIKey, data, client)
}
