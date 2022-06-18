package cmd

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/bounoable/deepl"
	"github.com/lucxjo/diru/cfg"
	"github.com/lucxjo/diru/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Commands is a command manager. It simply calls the appropriate function based on the command.
func Commands(msg *disgord.Message, s disgord.Session, c *deepl.Client, mdb *mongo.Collection, client *disgord.Client) {
	var guildPrefs cfg.GuildPrefs

	defaultSettings := cfg.GuildPrefs{
		GuildID:   msg.GuildID.String(),
		PreferredService: "deepl",
		DefaultLang: "en",
		DeepLEnabled: true,
		GtrEnabled: true,
	}
	err := mdb.FindOne(context.TODO(), bson.D{{"guildid", msg.GuildID.String()}}).Decode(&guildPrefs)

	if err == mongo.ErrNoDocuments {
		mdb.InsertOne(context.TODO(), defaultSettings)
		mdb.FindOne(context.TODO(), bson.D{{"guildid", msg.GuildID.String()}}).Decode(&guildPrefs)
	} else if err == nil {} else {
		fmt.Println("Database: " + err.Error())
	}

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
	} else if prefix == "toggle-provider" && cfg.GetValue("deepl_token").(string) != "" && cfg.GetValue("gtr_project_id").(string) != "" {
		m, err := client.Guild(msg.GuildID).Member(msg.Author.ID).GetPermissions()
		hasPermission := false

		if err != nil {
			msg.Reply(context.Background(), s, "Error: "+err.Error())
		}

		if m.Contains(disgord.PermissionManageServer) || m.Contains(disgord.PermissionAdministrator) { hasPermission = true }

		if hasPermission {
			funct := strings.Split(msg.Content, " ")[2]
			filter := bson.D{{"guildid", msg.GuildID.String()}}
			if funct == "default" {
				if guildPrefs.PreferredService == "deepl" {
					update := bson.D{{"$set", bson.D{{"preferredservice", "gtr"}}}}
					mdb.UpdateOne(context.TODO(), filter, update)
					msg.Reply(context.Background(), s, "Default provider set to Google Translate")
				} else {
					update := bson.D{{"$set", bson.D{{"preferredservice", "deepl"}}}}
					mdb.UpdateOne(context.TODO(), filter, update)
					msg.Reply(context.Background(), s, "Default provider set to DeepL")
				}
			}
		} else {
			msg.Reply(context.Background(), s, "You do not have permission to use this command.")
		}
	} else {
		if cfg.GetValue("deepl_token").(string) != "" && cfg.GetValue("gtr_project_id").(string) != "" {
			Dpla(msg, s, c)
		} else if cfg.GetValue("deepl_token").(string) != "" {
			Dpla(msg, s, c)
		} else {
			Gtra(msg, s)
		}
	}
}
