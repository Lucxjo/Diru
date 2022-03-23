package cmd

import (
	"context"
	"strings"

	"github.com/andersfylling/disgord"
	deeplgo "github.com/bounoable/deepl"
	"github.com/lucxjo/diru/deepl"
)

// Dpl uses the Translate function to translate text to a requested language.
func Dpl(msg *disgord.Message, s disgord.Session, c *deeplgo.Client) {
	sArr := strings.Split(msg.Content, " ")

	if len(sArr) <= 2 && msg.Type != disgord.MessageTypeReply {
		msg.Reply(context.Background(), s, "Please specify a language code.\n Please use the format `@Diru dpl <lang> <phrase>`.\n For language codes see: https://github.com/Lucxjo/Diru/wiki/Supported-translators-and-languages#deepls-supported-languages")
		return
	}

	lang := sArr[2]

	if !deepl.CheckCode(lang) {
		msg.Reply(context.Background(), s, "Language code is incorrect.\n Please use the format `@Diru dpl <lang> <phrase>`.\n For language codes see: https://github.com/Lucxjo/Diru/wiki/Supported-translators-and-languages#deepls-supported-languages")
		return
	}

	if msg.Type == disgord.MessageTypeReply {
		text := msg.ReferencedMessage.Content

		msg.Reply(context.Background(), s, deepl.TranslateTo(lang, text, c))
	} else {
		text := strings.Split(msg.Content, " ")[3:]
		tx := strings.Join(text, " ")

		msg.Reply(context.Background(), s, deepl.TranslateTo(lang, tx, c))
	}
}