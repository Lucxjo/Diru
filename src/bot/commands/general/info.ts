import {
	ActionRowBuilder,
	ButtonBuilder,
	ButtonStyle,
	CommandInteraction,
	EmbedBuilder,
} from 'discord.js';
import { Discord, SimpleCommand, SimpleCommandMessage, Slash } from 'discordx';
import pkgJson from '../../../../package.json' assert { type: 'json' };

@Discord()
export class Info {
	private memory = process.memoryUsage();
	private _embed = new EmbedBuilder()
		.setColor('Blurple')
		.setTitle('Bot information')
		.setDescription('Diru is a Discord bot that can translate text.')
		.addFields([
			{
				name: 'OS',
				value: process.platform,
				inline: true,
			},
			{
				name: 'Architecture',
				value: process.arch,
				inline: true,
			},
			{
				name: 'Version',
				value: pkgJson.version,
				inline: true,
			},
			{
				name: 'Memory usage',
				value: `${(
					this.memory.heapUsed / 1024 / 1024 +
					this.memory.heapTotal / 1024 / 1024
				).toFixed(2)} MB`,
				inline: true,
			},
		]);

	@Slash('info', {
		description: 'Get information about the bot.',
		dmPermission: true,
	})
	async slash(interaction: CommandInteraction) {
		const srcBtn = new ButtonBuilder()
			.setLabel('Source')
			.setEmoji('📦')
			.setURL('https://github.com/Lucxjo/Diru')
			.setStyle(ButtonStyle.Link);

		const issueBtn = new ButtonBuilder()
			.setLabel('Report an issue')
			.setEmoji('📩')
			.setStyle(ButtonStyle.Link)
			.setURL('https://github.com/Lucxjo/Diru/issues');

		const btnRow = new ActionRowBuilder<ButtonBuilder>().addComponents(
			srcBtn,
			issueBtn
		);
		interaction.reply({
			embeds: [this._embed],
			ephemeral: true,
			components: [btnRow],
		});
	}

	@SimpleCommand('info', { description: 'Get bot information' })
	async simple(cmd: SimpleCommandMessage) {
		cmd.message.reply({
			embeds: [this._embed],
		});
	}
}
