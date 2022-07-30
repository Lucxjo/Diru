import { ApplicationCommandType, MessageContextMenuCommandInteraction } from "discord.js";
import { ContextMenu, Discord, Guard } from "discordx";

@Discord()
export class Test {
	@ContextMenu(ApplicationCommandType.Message, "Test Interaction")
	async test(interaction: MessageContextMenuCommandInteraction) {
		console.info(interaction.targetMessage.content)
		interaction.reply({ content: `Test: ${interaction.targetMessage.content}`, ephemeral: true });
	}
}