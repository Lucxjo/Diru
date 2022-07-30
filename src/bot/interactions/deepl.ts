import { ApplicationCommandType, MessageContextMenuCommandInteraction } from "discord.js";
import { ContextMenu, Discord, Guard } from "discordx";
import deeplTranslate from "../helpers/deepl";

@Discord()
export class DeepL {
	@ContextMenu(ApplicationCommandType.Message, 'Send to DeepL (User)')
	async translateToLocale(interaction: MessageContextMenuCommandInteraction) {
		deeplTranslate(interaction.targetMessage.content, 'XX', interaction);
	}

	@ContextMenu(ApplicationCommandType.Message, 'Send to DeepL (Guild)')
	async translateToGuild(interaction: MessageContextMenuCommandInteraction) {
		console.info(interaction.targetMessage.content);
		deeplTranslate(
			interaction.targetMessage.content,
			interaction.guildLocale!.toUpperCase(),
			interaction
		);
	}

	@ContextMenu(ApplicationCommandType.Message, 'Send to DeepL (EN)')
	async translateToEnglish(interaction: MessageContextMenuCommandInteraction) {
		console.info(interaction.targetMessage.content);
		deeplTranslate(
			interaction.targetMessage.content,
			'EN-GB',
			interaction
		);
	}
}