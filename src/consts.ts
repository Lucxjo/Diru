import 'dotenv/config';
export const { DISCORD_TOKEN, CLIENT_ID, GUILD_ID } = process.env;

export const NODE_ENV = process.env.NODE_ENV ?? 'development';
export const MONGO_URI = process.env.MONGO_URI ?? 'mongodb://localhost:27017/diru';