package cmd

import (
	"context"
	"strings"

	"github.com/andersfylling/disgord"
	deeplgo "github.com/bounoable/deepl"
	"github.com/lucxjo/diru/deepl"
)

// Dpla uses the AutoTranslate function to translate text to English (British).
func Dpla(msg *disgord.Message, s disgord.Session, c *deeplgo.Client) {
	if msg.Type == disgord.MessageTypeReply {
		text := msg.ReferencedMessage.Content

		msg.Reply(context.Background(), s, deepl.AutoTranslate(text, c))
	} else if strings.Contains(msg.Content, "dpl") {
		text := strings.Split(msg.Content, " ")[2:]
		tx := strings.Join(text, " ")
		msg.Reply(context.Background(), s, deepl.AutoTranslate(tx, c))
	} else {
		// Because this is the option if no prefix is chosen, we need to start from index 1
		text := strings.Split(msg.Content, " ")[1:]
		tx := strings.Join(text, " ")
		msg.Reply(context.Background(), s, deepl.AutoTranslate(tx, c))
	}
}