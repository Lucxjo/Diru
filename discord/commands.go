package discord

import "github.com/andersfylling/disgord"

// Commands contains all the slash commands for the bot.
var Commands = []*disgord.CreateApplicationCommand{
	{
		Name: "deepl",
		Description: "Translates a phrase from a language into English (British) using Deepl",
		Options: []*disgord.ApplicationCommandOption{
			{
				Name: 	  "from_lang",
				Required: false,
				Type:    disgord.OptionTypeString,
				Description: "The language to translate from. Will try to automatically detect if not specified.",
				Choices: []*disgord.ApplicationCommandOptionChoice{
					{
						Name: "Bulgarian",
						Value: "BG",
					},
					{
						Name: "Czech",
						Value: "CS",
					},
					{
						Name: "Danish",
						Value: "DA",
					},
					{
						Name: "German",
						Value: "DE",
					},
					{
						Name: "Greek",
						Value: "EL",
					},
					{
						Name: "English",
						Value: "EN",
					},
					{
						Name: "Spanish",
						Value: "ES",
					},
					{
						Name: "Estonian",
						Value: "ET",
					},
					{
						Name: "Finnish",
						Value: "FI",
					},
					{
						Name: "French",
						Value: "FR",
					},
					{
						Name: "Hungarian",
						Value: "HU",
					},
					{
						Name: "Italian",
						Value: "IT",
					},
					{
						Name: "Japanese",
						Value: "JA",
					},
					{
						Name: "Lithuanian",
						Value: "LT",
					},
					{
						Name: "Latvian",
						Value: "LV",
					},
					{
						Name: "Dutch",
						Value: "NL",
					},
					{
						Name: "Polish",
						Value: "PL",
					},
					{
						Name: "Portuguese",
						Value: "PT",
					},
					{
						Name: "Romanian",
						Value: "RO",
					},
					{
						Name: "Russian",
						Value: "RU",
					},
					{
						Name: "Slovak",
						Value: "SK",
					},
					{
						Name: "Slovenian",
						Value: "SL",
					},
					{
						Name: "Swedish",
						Value: "SV",
					},
					{
						Name: "ZH",
						Value: "ZH",
					},
				},
			},
		},
	},
}