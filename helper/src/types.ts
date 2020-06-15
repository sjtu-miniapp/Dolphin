export interface UserAuth {
  openID: string;
  sessionID: string;
}

export type DataType = 'group' | 'task';

export interface CodeRecord {
  id: number;
  code: number;
  type: DataType;
  expirationTime: number; // timestamp
}
