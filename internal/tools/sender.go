package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type APIAnswer struct {
	Models []map[string]any `json:"data"`
}

func SendGetWithContext(url string, api_key string, client *http.Client) ([]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Ошибка при создании Get запроса: %s", err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", api_key))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при отправки Get запроса: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, err
	}
	var data APIAnswer
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Ошибка при декодировании ответа: %s", err)
	}
	models := make([]any, 0, len(data.Models))
	for idx := range data.Models {
		model := data.Models[idx]
		models = append(models, model["id"])
	}
	return models, nil
}
