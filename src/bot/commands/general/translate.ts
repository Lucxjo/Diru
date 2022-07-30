import axios from 'axios';
import {
	ApplicationCommandOptionType,
	CommandInteraction,
	EmbedBuilder,
	PermissionsBitField,
} from 'discord.js';
import { Discord, Slash, SlashGroup, SlashOption } from 'discordx';
import { SecureConnect } from '../../../shared/SecureConnect';

@Discord()
@SlashGroup({ name: 'translate', description: 'Translate text' })
@SlashGroup('translate')
export class Translate {
	@Slash('deepl', { description: 'Translate text using DeepL' })
	async deepl(
		@SlashOption('phrase', {
			description: 'The phrase to translate',
			type: ApplicationCommandOptionType.String,
			required: true,
		})
		phrase: string,
		@SlashOption('data', { description: 'This will show extra information about a translation', type: ApplicationCommandOptionType.Boolean, required: false })
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

		const embed = new EmbedBuilder()
			.setTitle('DeepL Translation')
			.setColor('Aqua')
		
		if (data) {
			embed.addFields(embedFields);
		} else {
			embed.setDescription(embedFields[2].value);
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
				embeds: [embed],
				ephemeral: false,
			});
		} else {
			interaction.reply({
				embeds: [embed],
				ephemeral: true,
			});
		}
	}
}
