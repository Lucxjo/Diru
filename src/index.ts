import 'reflect-metadata';
import { Server } from './server/main';
import { Bot } from './bot/main';
import { SecureConnect } from './shared/SecureConnect';

SecureConnect.generateKey();

Bot.start();
Server.start();
