import 'reflect-metadata';
import { SecureConnect } from '@lucxjo/diru-shared/SecureConnect';
import { Server } from './main';

SecureConnect.generateKey();
Server.start();
