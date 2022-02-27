package main

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	deeplgo "github.com/bounoable/deepl"
	"github.com/lucxjo/diru/cfg"
	"github.com/lucxjo/diru/cmd"
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

	client.Gateway().WithMiddleware(cont.HasBotMentionPrefix).MessageCreate(func(s disgord.Session, h *disgord.MessageCreate) {
		if !h.Message.Author.Bot {
			cmd.Commands(h.Message, s, dClient)
		}
		//h.Message.Reply(context.Background(), s, "For help, see https://github.com/lucxjo/diru/wiki")
	})

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	defer client.Gateway().StayConnectedUntilInterrupted()
}
