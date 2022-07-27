import { CommandInteraction, EmbedBuilder } from 'discord.js';
import { Discord, Slash } from 'discordx';

@Discord()
export class Test {
	@Slash('bot-test')
	async info(interaction: CommandInteraction) {
		interaction.reply('Hello, Diru!');
	}
}
