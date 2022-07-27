import { ActivityType } from 'discord.js';
import { ArgsOf, Client, Discord, On } from 'discordx';

@Discord()
export class Ready {
	@On('ready')
	async ready([]: ArgsOf<'ready'>, client: Client) {
		await client.guilds.fetch();
		await client.initApplicationCommands({
			global: { log: true },
			guild: { log: true },
		});

		console.info(`Logged in as ${client.user?.username}!`);
		console.info(
			`Invite link: https://discord.com/oauth2/authorize?client_id=${client.user?.id}&permissions=8&scope=bot%20applications.commands`
		);
		client.user?.setPresence({
			activities: [{ name: 'In development...', type: ActivityType.Watching }],
			status: 'online',
		});
	}
}