package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	deeplgo "github.com/bounoable/deepl"
	"github.com/joho/godotenv"
	"github.com/lucxjo/diru/deepl"
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

	deeplCont, _ := std.NewMsgFilter(context.Background(), client)
	gotransCont, _ := std.NewMsgFilter(context.Background(), client)

	deeplCont.SetPrefix("deepl")
	gotransCont.SetPrefix("gotrans")

	client.Gateway().WithMiddleware(deeplCont.HasPrefix).MessageCreate(func(s disgord.Session, h *disgord.MessageCreate) {

		if h.Message.Author.Bot {
			return
		} else if h.Message.Type == disgord.MessageTypeReply {
			reply := h.Message.ReferencedMessage.Content
			if strings.TrimLeft(h.Message.Content, "deepl") != "" {
				h.Message.Reply(context.Background(), s, deepl.TranslateTo(strings.TrimLeft(h.Message.Content, "deepl"), reply, dClient))
			} else {
				h.Message.Reply(context.Background(), s, deepl.AutoTranslate(reply, dClient))
			}
		} else {
			m := strings.TrimLeft(h.Message.Content, "deepl")
			translated := deepl.AutoTranslate(m, dClient)
			h.Message.Reply(context.Background(), s, translated)
		}
	})

	client.Gateway().WithMiddleware(gotransCont.HasPrefix).MessageCreate(func(s disgord.Session, h *disgord.MessageCreate) {
		h.Message.Author.SendMsgString(context.Background(), s, "Google Translate is currently not implemented")
	})

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	defer client.Gateway().StayConnectedUntilInterrupted()
}
