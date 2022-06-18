package cmd

import (
	"context"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/lucxjo/diru/google"
)

// Gtr uses the Translate function to translate text to a requested language.
func Gtra(msg *disgord.Message, s disgord.Session) {
	if msg.Type == disgord.MessageTypeReply {
		text := msg.ReferencedMessage.Content
		tr, _ := google.TranslateTo("en", text)

		msg.Reply(context.Background(), s, tr)
	} else if strings.Contains(msg.Content, "dpl") {
		text := strings.Split(msg.Content, " ")[2:]
		tx := strings.Join(text, " ")
		tr, _ := google.TranslateTo("en", tx)
		msg.Reply(context.Background(), s, tr)
	} else {
		// Because this is the option if no prefix is chosen, we need to start from index 1
		text := strings.Split(msg.Content, " ")[1:]
		tx := strings.Join(text, " ")
		tr, _ := google.TranslateTo("en", tx)
		msg.Reply(context.Background(), s, tr)
	}
}
