import 'reflect-metadata';
import { SecureConnect } from '@lucxjo/diru-shared/SecureConnect';
import { Bot } from './main';

SecureConnect.generateKey();
await Bot.start();
