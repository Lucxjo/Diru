import 'reflect-metadata';
import { SecureConnect } from 'diru-shared/SecureConnect';
import { Server } from './main';

SecureConnect.generateKey();
Server.start();
