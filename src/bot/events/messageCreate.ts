import { Discord, On, Client, ArgsOf } from 'discordx';

@Discord()
export class MessageCreate {
	@On('messageCreate')
	async onMessage([message]: ArgsOf<'messageCreate'>, client: Client) {
		if (Date.now() - Date.UTC(2023, 1, 28, 13, 0, 0) >= 0 && message.mentions.users.has(client.user?.id!)) {
			message.reply(
				`This type of command is no longer available. Please use the slash command instead.\n`
				+
				`Can't see the slash command? Get a staff member to reinvite the bot using the invite button in my profile or https://discord.com/oauth2/authorize?client_id=${client.user?.id!}&permissions=274878237760&scope=bot%20applications.commands`
			);
		} else {
			client.executeCommand(message);
		}
	}
}
