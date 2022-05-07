package cmd

import (
	"context"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/lucxjo/diru/google"
)

// Gtr uses the Translate function to translate text to a requested language.
func Gtr(msg *disgord.Message, s disgord.Session) {
	sArr := strings.Split(msg.Content, " ")

	if len(sArr) <= 2 && msg.Type != disgord.MessageTypeReply {
		msg.Reply(context.Background(), s, "Please specify a language code.\n Please use the format `@Diru dpl <lang> <phrase>`.\n For language codes see: https://github.com/Lucxjo/Diru/wiki/Supported-translators-and-languages#deepls-supported-languages")
		return
	}

	lang := sArr[2]

	if msg.Type == disgord.MessageTypeReply {
		text := msg.ReferencedMessage.Content

		tr, err := google.TranslateTo(lang, text)

		if err != nil {
			msg.Reply(context.Background(), s, "Error: "+err.Error())
			return
		}

		msg.Reply(context.Background(), s, tr)
	} else {
		text := strings.Split(msg.Content, " ")[3:]
		tx := strings.Join(text, " ")

		tr, err := google.TranslateTo(lang, tx)

		if err != nil {
			msg.Reply(context.Background(), s, "Error: "+err.Error())
			return
		}

		msg.Reply(context.Background(), s, tr)
	}
}