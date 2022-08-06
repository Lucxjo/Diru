import 'reflect-metadata';
import { Server } from 'diru-server/main';
import { Bot } from 'diru-discord/main';
import { SecureConnect } from './shared/SecureConnect';

SecureConnect.generateKey();

Bot.start();
Server.start();
