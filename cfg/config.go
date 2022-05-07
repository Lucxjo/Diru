package cfg

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/cristalhq/aconfig"
)

type Topgg struct {
	Token string `default:""`
	Id string `default:""`
}

type DiruConfig struct {
	DiscordToken string `default:""`
	DeeplToken   string `default:""`
	GtrToken     string `default:""`
	GtrProjectId string `default:""`
	Topgg Topgg
}

func GetConfig() DiruConfig {

	if _, err := os.Stat("config/Diru.json"); errors.Is(err, os.ErrNotExist) {
		ioutil.WriteFile("config/Diru.json", []byte("{\n    \"discord_token\": \"\",\n    \"deepl_token\": \"\"\n,\n    \"gtr_token\": \"\"\n}"), 0644)

		panic("config/Diru.json not found. It has been created for you, you must enter your values for discord_token, gtr_token, and gtr_project_id and deepl_token.")
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
		GtrToken:     cfg.GtrToken,
		GtrProjectId: cfg.GtrProjectId,
		Topgg: cfg.Topgg,
	}

	return config
}