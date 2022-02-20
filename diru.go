package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"github.com/joho/godotenv"
	"github.com/lucxjo/diru/deepl"
  	deeplgo "github.com/bounoable/deepl"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	client := disgord.New(disgord.Config{
		BotToken: os.Getenv("DISCORD_TOKEN"),
		ProjectName: "Diru",
		Intents: disgord.IntentGuildMessages | disgord.IntentDirectMessages,
	})

	if err != nil {
		panic(err)
	}

	dClient := deeplgo.New(os.Getenv("DEEPL_TOKEN"))

	deeplCont, _ := std.NewMsgFilter(context.Background(), client)
	gotransCont, _ := std.NewMsgFilter(context.Background(), client)

	deeplCont.SetPrefix("deepl")
	gotransCont.SetPrefix("gotrans")

	client.Gateway().WithMiddleware(deeplCont.HasPrefix).MessageCreate(func(s disgord.Session, h *disgord.MessageCreate) {
		m := strings.TrimLeft(h.Message.Content, "deepl")
		translated := deepl.AutoTranslate(m, dClient)
		h.Message.Author.SendMsgString(context.Background(), s, "Original: " + m + "\nTranslation: " + translated)
	})

	client.Gateway().WithMiddleware(gotransCont.HasPrefix).MessageCreate(func(s disgord.Session, h *disgord.MessageCreate) {
		h.Message.Author.SendMsgString(context.Background(), s, "Google Translate is currently not implemented")
	})

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	defer client.Gateway().StayConnectedUntilInterrupted()
}