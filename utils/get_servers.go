package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Guild struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Icon        string   `json:"icon"`
	Owner       bool     `json:"owner"`
	Permissions string   `json:"permissions"`
	Features    []string `json:"features"`
}

func GetConnectedServers(token string) []Guild {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, _ := http.NewRequest("GET", "https://discordapp.com/api/v9/users/@me/guilds", nil)
	req.Header.Add("Authorization", "Bot "+token)

	resp, err := httpClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var data []Guild

	guilds, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(guilds, &data)

	return data
}

func GetConnectedServerCount(token string) int {
	return len(GetConnectedServers(token))
}
