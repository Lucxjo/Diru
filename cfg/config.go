package cfg

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/cristalhq/aconfig"
)


type DiruConfig struct {
	DiscordToken string `default:""`
	DeeplToken   string `default:""`
}

func GetConfig() DiruConfig {

	if _, err := os.Stat("config/Diru.json"); errors.Is(err, os.ErrNotExist) {
		ioutil.WriteFile("config/Diru.json", []byte("{\n    \"discord_token\": \"\",\n    \"deepl_token\": \"\"\n}"), 0644)

		panic("config/Diru.json not found. It has been created for you, you must enter your values for discord_token and deepl_token.")
	}

	var cfg DiruConfig
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipDefaults: true,
		SkipEnv: true,
		SkipFlags: true,
		Files: []string{"config/Diru.json"},
	})

	if err := loader.Load(); err != nil {
		panic(err)
	}

	config := DiruConfig{
		DiscordToken: cfg.DiscordToken,
		DeeplToken:   cfg.DeeplToken,
	}

	return config
}