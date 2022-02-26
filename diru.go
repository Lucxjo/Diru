package main

import (
	"context"
	"fmt"
	"os"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	deeplgo "github.com/bounoable/deepl"
	"github.com/joho/godotenv"
	"github.com/lucxjo/diru/cmd"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	client := disgord.New(disgord.Config{
		BotToken:    os.Getenv("DISCORD_TOKEN"),
		ProjectName: "Diru",
		Intents:     disgord.IntentGuildMessages | disgord.IntentDirectMessages,
	})

	if err != nil {
		panic(err)
	}

	dClient := deeplgo.New(os.Getenv("DEEPL_TOKEN"))

	cont, _ := std.NewMsgFilter(context.Background(), client)

	client.Gateway().WithMiddleware(cont.HasBotMentionPrefix).MessageCreate(func(s disgord.Session, h *disgord.MessageCreate) {
		if (!h.Message.Author.Bot) {
			cmd.Commands(h.Message, s, dClient)
		}
		//h.Message.Reply(context.Background(), s, "For help, see https://github.com/lucxjo/diru/wiki")
	})

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	defer client.Gateway().StayConnectedUntilInterrupted()
}
