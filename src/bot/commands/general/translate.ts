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
import { autoLanguage, verifyLanguage } from '../../../bot/helpers/deepl';

@Discord()
@SlashGroup({ name: 'translate', description: 'Translate text' })
@SlashGroup('translate')
export class Translate {
	private static embed = new EmbedBuilder().setTitle('DeepL Translation');

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
		@SlashOption('target-language', {
			description: 'The language to translate to',
			type: ApplicationCommandOptionType.String,
			required: false,
		})
		targetLanguage: string = 'XX',
		interaction: CommandInteraction
	) {
		let embedFields: { name: string; value: string }[] = [
			{ name: 'Requested phrase:', value: phrase },
		];

		const { langCode, defaultCode } =
			targetLanguage === 'XX'
				? autoLanguage(interaction.locale)
				: verifyLanguage(targetLanguage);

		await axios
			.post('http://localhost:3000/api/translate/deepl', {
				text: phrase,
				KEY: SecureConnect.key,
				LANG_CODE: langCode === '' ? 'EN-GB' : langCode,
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
			if (!defaultCode) {
				Translate.embed
					.setDescription(embedFields[2].value)
					.setColor('Aqua');
			} else {
				Translate.embed
					.setDescription(
						`${embedFields[2].value}\n\n**Note:** The language supplied is not supported by DeepL, so the translation was done using British English.`
					)
					.setColor('Orange');
			}
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

	@SimpleCommand('dpl', {
		description: 'Translate text using DeepL',
		argSplitter: ', ',
	})
	async deeplSimple(
		@SimpleCommandOption('target-lang')
		targetLang: string,
		@SimpleCommandOption('phrase')
		phrase: string,
		command: SimpleCommandMessage
	) {
		if (!phrase) {
			command.message.reply({
				embeds: [
					new EmbedBuilder()
						.setTitle('Error')
						.setDescription(
							'You need to specify a phrase and a target language to translate\nUsage: `@Diru dpla <target-language>, <phrase>`\n\n**Note:** The comma and space is required between the target language and the phrase.'
						)
						.setColor('Red'),
				],
			});
			return;
		}

		let embedFields: { name: string; value: string }[] = [
			{ name: 'Requested phrase:', value: phrase },
		];
		const { langCode, defaultCode } = verifyLanguage(targetLang);

		await axios
			.post('http://localhost:3000/api/translate/deepl', {
				text: phrase,
				KEY: SecureConnect.key,
				LANG_CODE: langCode === '' ? 'EN-GB' : langCode,
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

		if (!defaultCode) {
			Translate.embed
				.setDescription(embedFields[2].value)
				.setColor('Aqua');
		} else {
			Translate.embed
				.setDescription(
					`${embedFields[2].value}\n\n**Note:** The language supplied is not supported by DeepL, so the translation was done using British English.`
				)
				.setColor('Orange');
		}

		command.message.reply({
			embeds: [Translate.embed],
		});
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
							'You need to specify a phrase to translate\nUsage: `@Diru dpla <phrase>`'
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

		Translate.embed.setDescription(embedFields[2].value).setColor('Aqua');

		command.message.reply({
			embeds: [Translate.embed],
		});
	}
}
