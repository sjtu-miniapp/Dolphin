import { UserAuth, CodeRecord } from '../types';
import { sharecode } from './conn';

export interface ShareCode {
  id: string;
  code: number;
  open_id: string;
  session_id: string;
  expiration_time: string;
}

export const insertNewRecord = async (r: UserAuth & CodeRecord): Promise<ShareCode[]> => {
  const sql = `INSERT INTO sharecode (id, code, open_id, session_id, expiration_time)
   VALUES ($(id), $(code), $(openID), $(sessionID), to_timestamp($(expirationTime) / 1000.0)) RETURNING *`;

  const params = {
    id: `${r.type}${r.id}`,
    code: r.code,
    openID: r.openID,
    sessionID: r.sessionID,
    expirationTime: r.expirationTime,
  };

  return sharecode.query<ShareCode[]>(sql, params);
};

export const queryRecord = async (code: number): Promise<ShareCode[]> => {
  return sharecode.query<ShareCode[]>('SELECT * FROM sharecode WHERE code = $(code)', { code });
};
