package cfg

import (
	"github.com/spf13/viper"
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

func Init(name string, defaults map[string]interface{}) {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}

	if name == "" {
		viper.SetConfigName("diru")
		viper.SetConfigName("Diru")
	} else {
		viper.SetConfigName(name)
	}

	viper.SetConfigType("toml")
	viper.SetConfigType("json")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
}

func GetValue(k string) interface{} {
	viper.ReadInConfig()
	return viper.Get(k)
}

func SetValue(k string, v interface{}) {
	panic("not implemented")
}