import 'reflect-metadata';
import { SecureConnect } from 'src/shared/SecureConnect';
import { Bot } from './main';

SecureConnect.generateKey();
await Bot.start();
