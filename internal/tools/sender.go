package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type APIAnswer struct {
	ID      string           `json:"id"`
	Choices []map[string]any `json:"choices"`
}

func SendPostWithContext(url string, api_key string, body_resp map[string]any, client *http.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	jsonData, err := json.Marshal(body_resp)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %s", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Ошибка при создании Get запроса: %s", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", api_key))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при отправки Get запроса: %s", err)
	}
	defer resp.Body.Close()
	var data APIAnswer
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatalf("Ошибка при декодировании: %s", err)
	}
	if m, ok := data.Choices[0]["message"].(map[string]any); ok {
		fmt.Printf("%v", m["content"])
	}
}
