import axios from 'axios';
import express from 'express';
import { DEEPL_TOKEN, NODE_ENV } from '../consts';
import { SecureConnect } from '../shared/SecureConnect';

export class Server {
	private static _app: express.Express;

	static get app(): express.Express {
		return this._app;
	}

	static start() {
		this._app = express();
		this._app.use(express.json());

		this.app.post(
			'/api/translate/deepl',
			async (req: express.Request, res: express.Response) => {
				const {
					text,
					KEY,
					LANG_CODE = 'EN-GB',
				}: {
					text: string | undefined;
					KEY: string | undefined;
					LANG_CODE: string;
				} = req.body;

				if (KEY) {
					if (KEY !== SecureConnect.key) {
						res.status(401).send('Invalid key!');
						return;
					}
				} else {
					res.status(401).send('No key!');
					return;
				}

				if (text) {
					const dataToSend = new URLSearchParams();
					dataToSend.append('text', text);
					dataToSend.append('target_lang', LANG_CODE);

					const translate = axios
						.post(
							`https://api.deepl.com/v2/translate?auth_key=${DEEPL_TOKEN}`,
							dataToSend
						)
						.then((response) => {
							console.log(response.data);
							res.send(response.data.translations[0]);
							return;
						})
						.catch((error) => {
							console.error(error);
							res.sendStatus(500).send('Translation error!');
							return;
						});
				} else {
					res.status(400).send('No text!');
				}
			}
		);

		this.app.get('/', (req, res) => {
			if (NODE_ENV === 'development') {
				res.send(
					`<h1>Hello Diru!</h1> <p>This is a test server.</p> <p>SecureConnect string: ${SecureConnect.key}</p>`
				);
			} else {
				res.send(`<h1>Hello Diru!</h1> <p>This is a test server.</p>`);
			}
		});

		this.app.listen(3000, () => {
			console.log('Server is running on port 3000');
		});
	}
}
