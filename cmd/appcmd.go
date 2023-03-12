package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andersfylling/disgord"
	"github.com/lucxjo/diru/cfg"
	"github.com/sirupsen/logrus"
)

type AppCmdType uint

const (
	// DEFAULT
	// Slash command
	CHAT_INPUT AppCmdType = 1

	// Application command that appears on the user profile
	USER AppCmdType = 2

	// Application command that appears in the context menu of a message
	MESSAGE AppCmdType = 3
)

func GetGlobalRegisteredAppCmds(c *disgord.Client) []disgord.ApplicationCommand {
	u, _ := c.CurrentUser().Get()
	id := u.ID
	logrus.Infof("ID: %v\n", id)
	httpClient := http.Client{}

	var cmds []disgord.ApplicationCommand

	req, err := http.NewRequest("GET", fmt.Sprintf(`https://discord.com/api/v10/applications/%v/commands`, id), nil)

	if err != nil {
		logrus.Errorf("Err making command GET request: \n%v\n", err)
		return nil
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bot %v", cfg.GetValue("discord_token").(string)))

	resp, err := httpClient.Do(req)

	if err != nil {
		logrus.Errorf("Err getting commands: \n%v\n", err)
		return nil
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&cmds)

	if err != nil {
		logrus.Errorf("Err decoding app cmd data: \n%v\n", err)
		return nil
	}

	//respData, err := ioutil.ReadAll(resp.Body)

	//if err != nil {
	//	logrus.Errorf("Err reading commands: \n%v\n", err)
	//	return
	//}

	fmt.Println(cmds)
	return cmds
}

func GetAllGuildRegisteredAppCmds(c *disgord.Client) {
	guilds, _ := c.CurrentUser().GetGuilds(&disgord.GetCurrentUserGuilds{})

	for _, guild := range guilds {
		GetGuildRegisteredAppCmds(c, guild.ID)
	}

}

func GetGuildRegisteredAppCmds(c *disgord.Client, gid disgord.Snowflake) []disgord.ApplicationCommand {
	u, _ := c.CurrentUser().Get()
	id := u.ID
	logrus.Infof("ID: %v\n", id)
	httpClient := http.Client{}
	var cmds []disgord.ApplicationCommand

	req, err := http.NewRequest("GET", fmt.Sprintf(`https://discord.com/api/v10/applications/%v/guilds/%v/commands`, id, gid), nil)

	if err != nil {
		logrus.Errorf("Err making command GET request: \n%v\n", err)
		return nil
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bot %v", cfg.GetValue("discord_token").(string)))

	resp, err := httpClient.Do(req)

	if err != nil {
		logrus.Errorf("Err getting commands: \n%v\n", err)
		return nil
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&cmds)

	if err != nil {
		logrus.Errorf("Err decoding app cmd data: \n%v\n", err)
		return nil
	}

	//respData, err := ioutil.ReadAll(resp.Body)

	//if err != nil {
	//	logrus.Errorf("Err reading commands: \n%v\n", err)
	//	return
	//}

	fmt.Println(cmds)

	return cmds
}
