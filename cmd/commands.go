package cmd

import (
	"context"
	"runtime"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/bounoable/deepl"
)

// Commands is a command manager. It simply calls the appropriate function based on the command.
func Commands(msg *disgord.Message, s disgord.Session, c *deepl.Client) {
	prefix := strings.Split(msg.Content, " ")[1]

	if prefix == "dpl" {
		Dpl(msg, s, c)
	} else if prefix == "dpla" {
		Dpla(msg, s, c)
	} else if prefix == "info" {
		msg.Reply(context.Background(), s, "Diru is a Discord bot that can translate text.\n\n**Technical information:**\n`" + "OS: " + runtime.GOOS + "`\n`" + "Arch: " + runtime.GOARCH + "`\n`" + "Go Version: " + runtime.Version()+"`\n`" + "Version: 1.0.0" + "`\n" + "Source: https://github.com/Lucxjo/Diru/")
	} else if prefix == "issue" {
		msg.Reply(context.Background(), s, "Please report any issues on the GitHub issue tracker: https://github.com/Lucxjo/Diru/issues")
	} else if prefix == "help" {
		msg.Reply(context.Background(), s, "**Commands**\nAll commands require the bot to be mentioned\n\n`@Diru dpl <lang> <phrase>`\nTranslates a phrase to a specified language.\n\n`@Diru dpla <phrase>`\nTranslates a phrase to English (British).\n\n`@Diru info`\nDisplays technical information about the bot.\n\n`@Diru issue`\nDisplays a link to the GitHub issue tracker.")
	} else {
		msg.Reply(context.Background(), s, "Command not found.\nPlease use `@Diru help` to see a list of available commands.")
	}
}