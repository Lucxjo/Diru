package main

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	deeplgo "github.com/bounoable/deepl"
	"github.com/lucxjo/diru/cfg"
	"github.com/lucxjo/diru/cmd"
	"github.com/lucxjo/diru/utils"

	"cloud.google.com/go/pubsub"
)

func main() {
	cfg.Init("", map[string]interface{}{
		"discord_token": "",
		"deepl_token":   "",
		"gtr_token":     "",
		"gtr_project_id": "",
		"topgg": map[string]interface{}{
			"token": "",
			"id":    "",
		},
	})

	client := disgord.New(disgord.Config{
		BotToken:    cfg.GetValue("discord_token").(string),
		ProjectName: "Diru",
		Intents:     disgord.IntentGuildMessages | disgord.IntentDirectMessages,
	})

	u, err := client.BotAuthorizeURL(disgord.PermissionUseSlashCommands, []string{"bot"})
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

	dClient := deeplgo.New(cfg.GetValue("deepl_token").(string))
	gClient, err := pubsub.NewClient(context.Background(), cfg.GetValue("gtr_project_id").(string))

	if err != nil {
		panic(err)
	}

	defer gClient.Close()

	cont, _ := std.NewMsgFilter(context.Background(), client)

	client.Gateway().BotReady(func() {
		fmt.Println("Bot is now running.  Press CTRL-C to exit.")
		bot, _ := client.Gateway().GetBot()

		if cfg.GetValue("topgg.token") != "" && cfg.GetValue("topgg.id") != "" {
			utils.SendTopggData(cfg.GetValue("topgg.token").(string), cfg.GetValue("topgg.id").(string), bot.Shards, cfg.GetValue("discord_token").(string))
		}
	})

	client.Gateway().WithMiddleware(cont.HasBotMentionPrefix).MessageCreate(func(s disgord.Session, h *disgord.MessageCreate) {
		if !h.Message.Author.Bot {
			cmd.Commands(h.Message, s, dClient)
		}
		//h.Message.Reply(context.Background(), s, "For help, see https://github.com/lucxjo/diru/wiki")
	})

	defer client.Gateway().StayConnectedUntilInterrupted()
}
