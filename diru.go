package main

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	deeplgo "github.com/bounoable/deepl"
	"github.com/lucxjo/diru/cfg"
	"github.com/lucxjo/diru/cmd"
	"github.com/lucxjo/diru/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"cloud.google.com/go/pubsub"
)

func main() {

	go utils.DiruHttp()

	cfg.Init("", map[string]interface{}{
		"discord_token":  "",
		"deepl_token":    "",
		"gtr_project_id": "",
		"db_uri":		 "mongodb://localhost:27017",
		"topgg": map[string]interface{}{
			"token": "",
			"id":    "",
		},
	})

	mClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.GetValue("db_uri").(string)))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	mcol := mClient.Database("diru").Collection("guildPrefs")

	var dClient *deeplgo.Client

	client := disgord.New(disgord.Config{
		BotToken:    cfg.GetValue("discord_token").(string),
		ProjectName: "Diru",
		Intents:     disgord.IntentGuildMessages | disgord.IntentDirectMessages,
	})

	u, err := client.BotAuthorizeURL(disgord.PermissionUseSlashCommands, []string{"bot"})
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
	if cfg.GetValue("deepl_token").(string) == "" && cfg.GetValue("gtr_project_id").(string) == "" {
		panic("No translation provider configured")
	}
	if cfg.GetValue("deepl_token").(string) != "" {
		dClient = deeplgo.New(cfg.GetValue("deepl_token").(string))
	}

	if cfg.GetValue("gtr_project_id").(string) != "" {
		gClient, err := pubsub.NewClient(context.Background(), cfg.GetValue("gtr_project_id").(string))

		if err != nil {
			panic(err)
		}

		defer gClient.Close()
	}

	cont, _ := std.NewMsgFilter(context.Background(), client)

	client.Gateway().BotReady(func() {
		fmt.Println("Bot is now running.  Press CTRL-C to exit.")
		bot, _ := client.Gateway().GetBot()

		if cfg.GetValue("topgg.token") != "" && cfg.GetValue("topgg.id") != "" {
			utils.SendTopggData(cfg.GetValue("topgg.token").(string), cfg.GetValue("topgg.id").(string), bot.Shards, cfg.GetValue("discord_token").(string))
		}
	})

	client.Gateway().WithMiddleware(cont.HasBotMentionPrefix).MessageCreate(func(s disgord.Session, h *disgord.MessageCreate) {
		if !h.Message.Author.Bot {
			cmd.Commands(h.Message, s, dClient, mcol, client)
		}
		//h.Message.Reply(context.Background(), s, "For help, see https://github.com/lucxjo/diru/wiki")
	})

	defer client.Gateway().StayConnectedUntilInterrupted()
}
