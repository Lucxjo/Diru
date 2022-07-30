import axios from 'axios';
import {
	ApplicationCommandOptionType,
	CommandInteraction,
	EmbedBuilder,
	PermissionsBitField,
} from 'discord.js';
import {
	Discord,
	SimpleCommand,
	SimpleCommandMessage,
	SimpleCommandOption,
	Slash,
	SlashGroup,
	SlashOption,
} from 'discordx';
import { SecureConnect } from '../../../shared/SecureConnect';

@Discord()
@SlashGroup({ name: 'translate', description: 'Translate text' })
@SlashGroup('translate')
export class Translate {
	private static embed = new EmbedBuilder()
		.setTitle('DeepL Translation')
		.setColor('Aqua');

	@Slash('deepl', { description: 'Translate text using DeepL' })
	async deeplSlash(
		@SlashOption('phrase', {
			description: 'The phrase to translate',
			type: ApplicationCommandOptionType.String,
			required: true,
		})
		phrase: string,
		@SlashOption('data', {
			description: 'This will show extra information about a translation',
			type: ApplicationCommandOptionType.Boolean,
			required: false,
		})
		data: boolean = false,
		interaction: CommandInteraction
	) {
		let embedFields: { name: string; value: string }[] = [
			{ name: 'Requested phrase:', value: phrase },
		];

		await axios
			.post('http://localhost:3000/api/translate/deepl', {
				text: phrase,
				KEY: SecureConnect.key,
			})
			.then((res) => {
				embedFields.push(
					{
						name: 'Detected language:',
						value: res.data.detected_source_language,
					},
					{
						name: 'Translated phrase:',
						value: res.data.text,
					}
				);
			});

		if (data) {
			Translate.embed.addFields(embedFields);
		} else {
			Translate.embed.setDescription(embedFields[2].value);
		}

		if (
			interaction.memberPermissions?.has(
				PermissionsBitField.Flags.Administrator
			) ||
			interaction.memberPermissions?.has(
				PermissionsBitField.Flags.ManageMessages
			)
		) {
			interaction.reply({
				embeds: [Translate.embed],
				ephemeral: false,
			});
		} else {
			interaction.reply({
				embeds: [Translate.embed],
				ephemeral: true,
			});
		}
	}

	@SimpleCommand('dpla', {
		description: 'Translate text to English (GB) using DeepL',
		argSplitter: ', ',
	})
	async deeplASimple(
		@SimpleCommandOption('phrase', {
			description: 'The phrase to be translated',
		})
		phrase: string,
		command: SimpleCommandMessage
	) {
		if (!phrase) {
			command.message.reply({
				embeds: [
					new EmbedBuilder()
						.setTitle('Error')
						.setDescription(
							'You need to specify a phrase to translate\nUsage: `@Diru dpl <language>, <phrase>`'
						)
						.setColor('Red'),
				],
			});
			return;
		}

		let embedFields: { name: string; value: string }[] = [
			{ name: 'Requested phrase:', value: phrase },
		];

		await axios
			.post('http://localhost:3000/api/translate/deepl', {
				text: phrase,
				KEY: SecureConnect.key,
			})
			.then((res) => {
				embedFields.push(
					{
						name: 'Detected language:',
						value: res.data.detected_source_language,
					},
					{
						name: 'Translated phrase:',
						value: res.data.text,
					}
				);
			});

		Translate.embed.setDescription(embedFields[2].value);

		command.message.reply({
			embeds: [Translate.embed],
		});
	}
}
