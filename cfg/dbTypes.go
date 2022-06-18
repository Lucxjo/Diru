package cfg

type GuildPrefs struct {
	GuildID   string `json:"gid"`
	PreferredService string `json:"preferredService"`
	DefaultLang string `json:"defaultLang"`
	DeepLEnabled bool `json:"deepLEnabled"`
	GtrEnabled bool `json:"gtrEnabled"`
}