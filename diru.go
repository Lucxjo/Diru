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
)

func main() {
	config := cfg.GetConfig()

	client := disgord.New(disgord.Config{
		BotToken:    config.DiscordToken,
		ProjectName: "Diru",
		Intents:     disgord.IntentGuildMessages | disgord.IntentDirectMessages,
	})

	dClient := deeplgo.New(config.DeeplToken)

	cont, _ := std.NewMsgFilter(context.Background(), client)

	client.Gateway().BotReady(func() {
		fmt.Println("Bot is now running.  Press CTRL-C to exit.")
		bot, _ := client.Gateway().GetBot()

		if config.Topgg.Token != "" && config.Topgg.Id != "" {
			utils.SendTopggData(config.Topgg.Token, config.Topgg.Id, bot.Shards)
		}
	})

	client.Gateway().WithMiddleware(cont.HasBotMentionPrefix).MessageCreate(func(s disgord.Session, h *disgord.MessageCreate) {
		if !h.Message.Author.Bot {
			cmd.Commands(h.Message, s, dClient, config)
		}
		//h.Message.Reply(context.Background(), s, "For help, see https://github.com/lucxjo/diru/wiki")
	})

	defer client.Gateway().StayConnectedUntilInterrupted()
}
