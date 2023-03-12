package utils

import (
	"io/ioutil"
	"net/http"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/snowflake/v5"
)

type DisClient struct {
	Token       string
	ID          string
	ProjectName string
	Intents     disgord.Intent
	Logger      disgord.Logger

	// Disgord client
	C *disgord.Client
}

type DAppCmd struct {
	Id          snowflake.Snowflake                 `json:"id"`
	Name        string                              `json:"name"`
	NameLoc     []interface{}                       `json:"name_localizations"`
	Type        int                                 `json:"type"`
	Desc        string                              `json:"description"`
	Options     []*disgord.ApplicationCommandOption `json:"options"`
	DefPerm     bool                                `json:"default_permission"`
	DefMemPerms []interface{}                       `json:"default_member_permissions"`
	Nsfw        bool                                `json:"nsfw"`
	Version     snowflake.Snowflake                 `json:"version"`
}

func (d *DisClient) CreateDisgordClient() {
	d.C = disgord.New(disgord.Config{
		BotToken:    d.Token,
		Intents:     d.Intents,
		ProjectName: d.ProjectName,
		Logger:      d.Logger,
	})
}

func (d *DisClient) InitFunctionParams() {
	user, err := d.C.CurrentUser().Get()

	if err != nil {
		d.Logger.Error("An error occurred: ", err)
	}

	d.ID = user.ID.String()
}

// Sends a get request to the Discord api endpoints with authorization.
// Endpoint must be a string without the full url or a leading `/`
func (d DisClient) HttpGet(endpoint string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://discord.com/api/v10/"+endpoint, nil)

	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bot "+d.Token)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	resBody, _ := ioutil.ReadAll(resp.Body)

	return string(resBody), nil

}
