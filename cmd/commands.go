package cmd

import (
	"context"
	"runtime"
	"strconv"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/bounoable/deepl"
	"github.com/lucxjo/diru/cfg"
	"github.com/lucxjo/diru/utils"
)

// Commands is a command manager. It simply calls the appropriate function based on the command.
func Commands(msg *disgord.Message, s disgord.Session, c *deepl.Client) {
	var mStats runtime.MemStats
	prefix := strings.Split(msg.Content, " ")[1]

	if prefix == "dpl" && cfg.GetValue("deepl_token").(string) != "" {
		Dpl(msg, s, c)
	} else if prefix == "dpla" && cfg.GetValue("deepl_token").(string) != "" {
		Dpla(msg, s, c)
	} else if prefix == "gtra" && cfg.GetValue("gtr_project_id").(string) != "" {
		Gtra(msg, s)
	} else if prefix == "gtr" && cfg.GetValue("gtr_project_id").(string) != "" {
		Gtr(msg, s)
	} else if prefix == "info" {
		runtime.ReadMemStats(&mStats)
		stats := strconv.FormatInt(int64(mStats.HeapInuse/1024/1024+mStats.HeapIdle/1024/1024), 10)

		msg.Reply(context.Background(), s, "Diru is a Discord bot that can translate text.\n\n**Technical information:**\n```"+
			"OS: "+runtime.GOOS+"\n"+"Arch: "+runtime.GOARCH+"\n"+
			"Version: 1.2.0"+"\n"+"Source: https://github.com/Lucxjo/Diru/\n"+"Usage: "+stats+" MB```")
	} else if prefix == "issue" {
		msg.Reply(context.Background(), s, "Please report any issues on the GitHub issue tracker: https://github.com/Lucxjo/Diru/issues")
	} else if prefix == "vote" && cfg.GetValue("topgg.token").(string) != "" && cfg.GetValue("topgg.id").(string) != "" {
		msg.Reply(context.Background(), s, "Vote for Diru on top.gg: https://top.gg/bot/"+cfg.GetValue("topgg.id").(string)+"/vote")
	} else if prefix == "help" {
		bot, _ := s.Gateway().GetBot()
		utils.SendTopggData(cfg.GetValue("topgg.token").(string), cfg.GetValue("topgg.id").(string), bot.Shards, cfg.GetValue("discord_token").(string))
		dplInfo := ""
		gtrInfo := ""

		if cfg.GetValue("deepl_token").(string) != "" {
			dplInfo = "`@Diru dpl <lang> <phrase>`\nTranslates a phrase to a specified language with DeepL.\n\n" +
				"`@Diru dpla <phrase>`\nTranslates a phrase to English (British) with DeepL.\n\n" +
				"`@Diru <phrase>` \ndoes the same thing as `dpla` or `gtra` (depending on config)\n\n"
		}

		if cfg.GetValue("gtr_project_id").(string) != "" {
			gtrInfo = "`@Diru gtr <lang> <phrase>`\nTranslates a phrase to a specified language with Google Translate\n\n" +
				"`@Diru gtra <phrase>`\nTranslates a phrase to English  with Google Translate\n\n"
		}

		msg.Reply(context.Background(), s, "**Commands**\nAll commands require the bot to be mentioned\n\n"+
			dplInfo+
			gtrInfo+
			"`@Diru info`\nDisplays technical information about the bot.\n\n"+
			"`@Diru issue`\nDisplays a link to the GitHub issue tracker.")
	} else {
		// Default to DPLA if none of the above prefixes and the token exists
		if cfg.GetValue("deepl_token").(string) != "" {
			Dpla(msg, s, c)
		} else {
			Gtra(msg, s)
		}
	}
}
