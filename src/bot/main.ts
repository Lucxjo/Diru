import { Client } from 'discordx';
import { importx, dirname } from '@discordx/importer';
import {
	GUILD_ID,
	NODE_ENV,
	DISCORD_TOKEN,
	DISCORD_CLIENT_ID,
} from 'diru-shared/consts';
import { IntentsBitField } from 'discord.js';
import { SecureConnect } from 'diru-shared/SecureConnect';

export class Bot {
	private static _client: Client;

	static get client(): Client {
		return this._client;
	}

	static async start() {
		if (DISCORD_CLIENT_ID === undefined || DISCORD_TOKEN === undefined) {
			throw new Error(
				'DISCORD_CLIENT_ID and DISCORD_TOKEN must be defined'
			);
		}

		console.info('Starting bot...');

		await importx(
			dirname(import.meta.url) +
				'/{commands,events,interactions}/**/*.{ts,js}'
		);

		if (NODE_ENV === 'development') {
			console.log(`Discord: ${SecureConnect.key}`);
		}

		console.log(`ENV: ${NODE_ENV}`);
		this._client = new Client({
			simpleCommand: {
				prefix: [`<@${DISCORD_CLIENT_ID}>`, `<@!${DISCORD_CLIENT_ID}>`],
			},
			intents: [
				IntentsBitField.Flags.Guilds,
				IntentsBitField.Flags.GuildMessages,
				IntentsBitField.Flags.DirectMessages,
				IntentsBitField.Flags.MessageContent,
			],
			botGuilds:
				NODE_ENV === 'development' ? [GUILD_ID ?? ''] : undefined,
		});

		console.info(this.client.botGuilds);

		if (!DISCORD_TOKEN) {
			throw new Error('DISCORD_TOKEN is not set!');
		} else await this._client.login(DISCORD_TOKEN!);
	}
}
