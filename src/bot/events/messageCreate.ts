import { Discord, On, Client, ArgsOf } from 'discordx';

@Discord()
export class MessageCreate {
	@On('messageCreate')
	async onMessage([message]: ArgsOf<'messageCreate'>, client: Client) {
		client.executeCommand(message);
	}
}
