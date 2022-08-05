import fastify, { FastifyInstance } from 'fastify';
import { DEEPL_TOKEN, NODE_ENV } from '../consts';
import { SecureConnect } from '../shared/SecureConnect';
import pug from 'pug';
import { TypeBoxTypeProvider } from '@fastify/type-provider-typebox';
import { DeepLType } from './types/deepl';
import console from 'console';

export class Server {
	static start() {
		const app = fastify().withTypeProvider<TypeBoxTypeProvider>();

		app.register(import('@fastify/view'), {
			engine: {
				pug: pug,
			},
		});

		const pages: { route: string; name: string; data?: object }[] = [
			{
				route: '/',
				name: 'index.dev',
			},
			{
				route: '/',
				name: 'index',
			},
		];

		pages.forEach((page) => {
			pug.compileFile(`./src/server/pages/${page.name}.pug`);
		});

		app.post<{ Body: DeepLType }>(
			'/api/translate/deepl',
			async (request, reply) => {
				const { text, KEY, LANG_CODE = 'EN-GB' } = request.body;

				if (KEY) {
					if (KEY !== SecureConnect.key) {
						reply.status(403).send('Invalid key!');
						return;
					}
				} else {
					reply.status(403).send('No key!');
					return;
				}

				if (text) {
					await fetch(
						`https://api.deepl.com/v2/translate?text=${encodeURIComponent(
							text
						)}&target_lang=${encodeURIComponent(LANG_CODE)}`,
						{
							method: 'POST',
							headers: {
								Authorization: `DeepL-Auth-Key ${DEEPL_TOKEN}`,
							},
						}
					)
						.then((response) => response.json())
						.then((data) => {
							console.log(data);
							reply.status(200).send(data.translations[0]);
						})
						.catch((error) => {
							console.log(error);
						});
				} else {
					reply.status(400).send('No text!');
				}
			}
		);

		app.get('/ping', async (request, reply) => {
			reply.type('application/json').send({ apiRunning: true });
		});

		app.get('/', (request, reply) => {
			if (NODE_ENV === 'development') {
				reply.view('./src/server/pages/index.dev.pug', {
					scString: SecureConnect.key,
				});
			} else {
				reply.view('./src/server/pages/index.pug');
			}
		});

		app.listen({ port: 3000 }, (err, address) => {
			if (err) {
				console.error(err);
				process.exit(1);
			}
			console.log(`Server listening on ${address}`);
		});
	}
}
