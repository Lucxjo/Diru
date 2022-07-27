import express from 'express';


export class Server {
	private static _app: express.Express;

	static get app(): express.Express {
		return this._app;
	}

	static start() {
		this._app = express();

		this.app.get('/', (req, res) => {
			res.send('Hello Diru!');
		});

		this.app.listen(3000, () => {
			console.log('Server is running on port 3000');
		});
	}
}