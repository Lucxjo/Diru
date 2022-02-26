package cmd

import (
	"context"
	"runtime"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/bounoable/deepl"
)

func Commands(msg *disgord.Message, s disgord.Session, c *deepl.Client) {
	prefix := strings.Split(msg.Content, " ")[1]

	if prefix == "dpl" {
		Dpl(msg, s, c)
	} else if prefix == "dpla" {
		Dpla(msg, s, c)
	} else if prefix == "info" {
		msg.Reply(context.Background(), s, "Diru is a Discord bot that can translate text.\n\n**Technical information:**\n`" + "OS: " + runtime.GOOS + "`\n`" + "Arch: " + runtime.GOARCH + "`\n`" + "Go Version: " + runtime.Version()+"`\n`" + "Version: 1.0.0" + "`\n`" + "Source: https://github.com/Lucxjo/Diru/`")
	}
}