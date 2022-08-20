import {
	ApplicationCommandOptionType,
	CommandInteraction,
	EmbedBuilder,
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
import deeplTranslate from '../../../bot/helpers/deepl';

@Discord()
@SlashGroup({ name: 'translate', description: 'Translate text' })
@SlashGroup('translate')
export class Translate {
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
		deeplTranslate(phrase, targetLanguage, interaction, data);
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
							'You need to specify a phrase and a target language to translate\nUsage: `@Diru dpl <target-language>, <phrase>`\n\n**Note:** The comma and space is required between the target language and the phrase.\n\nIf you are trying to translate a message, read the documentation on the changes between Diru.Go and Diru.Node [here](https://github.com/Lucxjo/Diru/wiki/Differences-between-Diru.Go-and-Diru.Node)'
						)
						.setColor('Red'),
				],
			});
			return;
		}

		deeplTranslate(phrase, targetLang, command);
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
							'You need to specify a phrase to translate\nUsage: `@Diru dpla <phrase>`\n\nIf you are trying to translate a message, read the documentation on the changes between Diru.Go and Diru.Node [here](https://github.com/Lucxjo/Diru/wiki/Differences-between-Diru.Go-and-Diru.Node)'
						)
						.setColor('Red'),
				],
			});
			return;
		}

		deeplTranslate(phrase, 'EN-GB', command);
	}
}
