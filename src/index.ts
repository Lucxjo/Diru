import 'reflect-metadata';
import { Server } from '@lucxjo/diru-server/main';
import { Bot } from '@lucxjo/diru-discord/main';
import { SecureConnect } from '@lucxjo/diru-shared/SecureConnect';

SecureConnect.generateKey();

Bot.start();
Server.start();
