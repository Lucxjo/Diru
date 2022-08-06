/**
 * Ensures that only the bot can access the API.
 * This is more of a preventative measure for if the API is exposed for whatever reason.
 * The API will need to be exposed for the Bot website to work.
 */
export class SecureConnect {
	private static _key: string;

	static get key(): string {
		return this._key;
	}

	static generateKey() {
		const rand = (length = 8) => {
			let chars =
				'ABCDEFGHIJKLMNÑOPQRSTUVWXYZÄÅÖabcdefghijklmnñopqrstuvwxyzäåö0123456789!?@#$%^&*()_+-=[]{}|;:,./£€|';
			let str = '';
			for (let i = 0; i < length; i++) {
				str += chars.charAt(Math.floor(Math.random() * chars.length));
			}

			return str;
		};

		this._key = rand(64);
	}
}
