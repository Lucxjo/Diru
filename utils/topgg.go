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
		Timeout: 60 * time.Second,
	}

	reqBody, err := json.Marshal(map[string]uint{
		"server_count": uint(GetConnectedServerCount(dctoken)),
		"shard_count": shards,
	})

	if err != nil {
		fmt.Println("Error marshalling top.gg data:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://top.gg/api/bots/" + botId + "/stats", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	if err != nil {
		fmt.Println("Error starting request, top.gg data:", err)
		return
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
}
