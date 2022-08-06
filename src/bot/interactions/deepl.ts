import {
	ActionRowBuilder,
	ApplicationCommandType,
	MessageContextMenuCommandInteraction,
	ModalActionRowComponentBuilder,
	ModalBuilder,
	ModalSubmitInteraction,
	TextInputBuilder,
	TextInputStyle,
} from 'discord.js';
import { ComponentType, ContextMenu, Discord, ModalComponent } from 'discordx';
import deeplTranslate from '../helpers/deepl';

@Discord()
export class DeepL {
	@ContextMenu(ApplicationCommandType.Message, 'Translate with DeepL')
	async translateToLocale(interaction: MessageContextMenuCommandInteraction) {
		let modal = new ModalBuilder()
			.setTitle('Translate with DeepL')
			.setCustomId('translate-deepl');

		const phraseInput = new TextInputBuilder()
			.setLabel('Phrase to translate')
			.setCustomId('phrase-deepl')
			.setStyle(TextInputStyle.Paragraph)
			.setValue(interaction.targetMessage.content);

		const langInput = new TextInputBuilder()
			.setLabel('Target Language')
			.setPlaceholder(
				'Enter a language or leave blank to use the default'
			)
			.setCustomId('target-lang-deepl')
			.setRequired(false)
			.setStyle(TextInputStyle.Short);

		const actionRows = [
			new ActionRowBuilder<ModalActionRowComponentBuilder>().addComponents(
				phraseInput
			),
			new ActionRowBuilder<ModalActionRowComponentBuilder>().addComponents(
				langInput
			),
		];

		actionRows.forEach((row) => {
			modal.addComponents(row);
		});

		await interaction.showModal(modal);
		//deeplTranslate(interaction.targetMessage.content, 'XX', interaction);
	}

	@ModalComponent('translate-deepl')
	async handleModal(interaction: ModalSubmitInteraction) {
		const [phrase, targetLang] = ['phrase-deepl', 'target-lang-deepl'].map(
			(id) => interaction.fields.getTextInputValue(id)
		);

		deeplTranslate(phrase, targetLang === '' ? 'XX' : targetLang, interaction);
	}
}
