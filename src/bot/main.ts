import { Client } from 'discordx';
import { importx, dirname } from '@discordx/importer';
import { GUILD_ID, NODE_ENV, DISCORD_TOKEN } from '../consts';
import { IntentsBitField } from 'discord.js';

export class Bot {
	private static _client: Client;

	static get client(): Client {
		return this._client;
	}

	static async start() {
		console.info('Starting bot...');

		await importx(
			dirname(import.meta.url) + '/{commands,events,interactions}/**/*.{ts,js}'
		);

		console.log(`ENV: ${NODE_ENV}`);
		this._client = new Client({
			simpleCommand: {
				prefix: ['<@946066668734545971>', '<@!946066668734545971>'],
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
