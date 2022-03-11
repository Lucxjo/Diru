package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func SendTopggData(token string, botId string, shards uint, dctoken string) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	reqBody, _ := json.Marshal(map[string]uint{
		"server_count": uint(GetConnectedServerCount(dctoken)),
		"shard_count": shards,
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
