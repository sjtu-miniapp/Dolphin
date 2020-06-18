// Temporary store data

import { UserAuth, CodeRecord } from './types';

const CODE_RECORD: (CodeRecord & UserAuth)[] = [];

export const updateRecord = (p: CodeRecord & UserAuth): number => {
  const { id, type } = p;

  // Check if target group or task already exist
  const existRecord = CODE_RECORD.filter((r) => r.id === id && r.type === type);

  // find one record, refresh whatever
  if (existRecord.length === 1) {
    const idx = CODE_RECORD.findIndex((r) => r.type === type && r.id === id);
    if (idx < 0 || idx >= CODE_RECORD.length) {
      throw new Error('Failed to find record with previous verified');
    }

    CODE_RECORD[idx] = { ...p };
    return CODE_RECORD[idx].code;
  }

  // find more than one record, delete all of them
  if (existRecord.length > 1) {
    // delete them
    for (let i = CODE_RECORD.length - 1; i >= 0; i--) {
      if (CODE_RECORD[i].id === id && CODE_RECORD[i].type === type) {
        CODE_RECORD.splice(i, 1);
      }
    }
  }

  // no record, add new record
  CODE_RECORD.push({ ...p });
  return CODE_RECORD[CODE_RECORD.length - 1].code;
};

export const getRecord = (code: number): (CodeRecord & UserAuth) | null => {
  console.log(1111, code, typeof code);
  console.log(2222, JSON.stringify(CODE_RECORD));
  const existRecord = CODE_RECORD.filter((r) => r.code === code);
  console.log(JSON.stringify(existRecord));

  if (existRecord.length > 1 || existRecord.length === 0) return null;

  return existRecord[0];
};
