package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TopggBotData struct {
	ServerCount uint `json:"server_count"`
	ShardCount uint `json:"shard_count"`
}

func SendTopggData(token string, botId string, shards uint) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	reqBody, _ := json.Marshal(TopggBotData{
		ServerCount: uint(GetConnectedServerCount(token)),
		ShardCount: shards,
	})

	req, _ := http.NewRequest("POST", "https://top.gg/api/bots/" + botId + "/stats", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	resp, err := httpClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
}
