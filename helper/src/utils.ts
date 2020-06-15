import * as moment from 'moment';

import { UserAuth, CodeRecord } from './types';
import { OPEN_ID, SESSION_ID } from './constants';
import * as store from './store';

export type ParsedQuery = Record<string, any> & UserAuth;

export const queryParser = (query: any): ParsedQuery => {
  if (!query) throw new Error('Invalid query');

  if (!query[OPEN_ID] || typeof query[OPEN_ID] !== 'string') throw new Error('Invalid open id');
  if (!query[SESSION_ID] || typeof query[SESSION_ID] !== 'string') throw new Error('Invalid session id');

  return query;
};

// Generate 6-digit number
export const codeGenerator = (): number => {
  return Math.floor(100000 + Math.random() * 900000);
};

export const expirationDateGenerator = (): number => {
  return moment().add(30, 'm').valueOf();
};

export const validateCode = (code: any): (CodeRecord & UserAuth) | null => {
  if (!code || isNaN(code)) return null;

  const parsedCode = parseInt(code, 10);

  if (parsedCode < 100000 || parsedCode > 999999) return null;

  return store.getRecord(parsedCode);
};
