import {
	CommandInteraction,
	EmbedBuilder,
	MessageContextMenuCommandInteraction,
	PermissionsBitField,
} from 'discord.js';
import { SimpleCommandMessage } from 'discordx';
import { SecureConnect } from '../../shared/SecureConnect';
import deeplLangs from '../../shared/deepl.json' assert { type: 'json' };

type DeepLTranslateOptions = typeof deeplLangs;
const deepLTranslateOptions: DeepLTranslateOptions = deeplLangs;

export const autoLanguage = (usrLocale: string) => {
	let langCode: string = '';
	let defaultCode = false;

	if (usrLocale.includes('-')) {
		let _langCode: string[] | string = usrLocale.split('-');
		if (_langCode[0].toUpperCase() === 'EN') {
			if (_langCode[1].toUpperCase() === 'US') {
				langCode = 'EN-US';
			} else {
				langCode = 'EN-GB';
			}
		} else if (_langCode[0].toUpperCase() === 'PT') {
			if (_langCode[1].toUpperCase() === 'BR') {
				langCode = 'PT-BR';
			} else {
				langCode = 'PT-PT';
			}
		} else {
			_langCode = _langCode[0].toUpperCase();
			deepLTranslateOptions.target.forEach((lang) => {
				if (lang.code === _langCode) {
					langCode = lang.code;
				}
			});
			if (!langCode) {
				langCode = 'EN-GB';
				defaultCode = true;
			}
		}
	} else {
		if (usrLocale.toUpperCase() === 'EN') {
			langCode = 'EN-GB';
		} else if (usrLocale.toUpperCase() === 'PT') {
			langCode = 'PT-PT';
		} else {
			deepLTranslateOptions.target.forEach((lang) => {
				if (lang.code.toUpperCase() === usrLocale.toUpperCase()) {
					langCode = lang.code;
				}
			});
			if (!langCode) {
				langCode = 'EN-GB';
				defaultCode = true;
			}
		}
	}

	return { langCode, defaultCode };
};

export const verifyLanguage = (lang: string) => {
	let l: string = '';
	let langCode: string | null = null;
	let defaultCode = false;
	if (lang.includes('-') || lang.includes('_')) {
		let _langCode = lang.toUpperCase();
		if (_langCode.includes('_')) {
			_langCode = _langCode.replace('_', '-');
		}
		({ langCode, defaultCode } = autoLanguage(_langCode));
	} else if (lang.length === 2) {
		({ langCode, defaultCode } = autoLanguage(lang.toUpperCase()));
	} else {
		deepLTranslateOptions.target.forEach(({ name, code, aliases }) => {
			if (name.toUpperCase() === lang.toUpperCase()) {
				l = code;
			} else if (aliases) {
				aliases.forEach((alias) => {
					if (alias.toUpperCase() === lang.toUpperCase()) {
						l = code;
					}
				});
			}
		});
		console.log(`bot/helpers/deepl.ts:62 ${l}`);
		({ langCode, defaultCode } = autoLanguage(l));
	}

	return { langCode, defaultCode };
};

const deeplTranslate = async (
	phrase: string,
	targetLanguage = 'XX',
	interaction:
		| MessageContextMenuCommandInteraction
		| CommandInteraction
		| SimpleCommandMessage,
	data = false
) => {
	let embed = new EmbedBuilder().setTitle('DeepL Translation');
	let embedFields: { name: string; value: string }[] = [
		{ name: 'Requested phrase:', value: phrase },
	];

	let descEnd = '';

	if (interaction instanceof SimpleCommandMessage)
		descEnd =
			'\n\nSomething not looking quite right? Try the slash command instead.';

	const { langCode, defaultCode } =
		targetLanguage === 'XX' &&
		!(interaction instanceof SimpleCommandMessage)
			? autoLanguage(interaction.locale)
			: verifyLanguage(targetLanguage);

	await fetch('http://localhost:3000/api/translate/deepl', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({
			text: phrase,
			LANG_CODE: langCode === '' ? 'EN-GB' : langCode,
			KEY: SecureConnect.key,
		}),
	})
		.then(async (res) => {
			let data = await res.json();
			console.log(`bot/helpers/deepl.ts:86 ${JSON.stringify(data)}`);
			embedFields.push(
				{
					name: 'Detected language:',
					value: data.detected_source_language,
				},
				{
					name: 'Translated phrase:',
					value: data.text,
				}
			);
		})
		.catch(console.error);

	if (!defaultCode) {
		embed.setColor('Aqua');
		if (data) embed.addFields(embedFields);
		else embed.setDescription(embedFields[2].value + descEnd);
	} else {
		embed.setColor('Orange');
		embed.setDescription(
			(data ? '' : embedFields[2].value + '\n\n') +
				'**Note:** The language supplied is not supported by DeepL, so the translation was done using British English.' +
				descEnd
		);
		if (data) embed.addFields(embedFields);
	}

	if (interaction instanceof CommandInteraction) {
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
	} else if (interaction instanceof MessageContextMenuCommandInteraction) {
		interaction.reply({ embeds: [embed] });
	} else {
		interaction.message.reply({ embeds: [embed] });
	}
};

export default deeplTranslate;
