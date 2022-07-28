import { ArgsOf, Client, Discord, On } from "discordx";

@Discord()
export class Interaction {
	@On('interactionCreate')
	async interactionCreate([interaction]: ArgsOf<'interactionCreate'>, client: Client) {
		client.executeInteraction(interaction);
	}
}