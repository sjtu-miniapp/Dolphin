import * as Router from '@koa/router';

import * as store from './store';
import * as utils from './utils';
import { CodeRecord, UserAuth } from './types';

const router = new Router({
  prefix: '/api/internal/v1',
});

router.get('/code', (ctx) => {
  const query = utils.queryParser(ctx.request.query);

  const { type, id, openID, sessionID } = query;

  if (type !== 'group' && type !== 'task') throw new Error('Invalid type');

  if (!id || isNaN(id)) throw new Error('Invaild id');

  const codeRecord: CodeRecord & UserAuth = {
    id,
    type,
    openID,
    sessionID,
    code: utils.codeGenerator(),
    expirationTime: utils.expirationDateGenerator(),
  };

  const addCodeResult = store.updateRecord(codeRecord);

  ctx.response.body = addCodeResult;
  ctx.response.status = 200;
});

router.put('/code/:code', (ctx) => {
  const query = utils.queryParser(ctx.request.query);
  console.log(query, ctx.params);

  const { code } = ctx.params;

  // TODO:
  // 1. find code in the db
  const record = utils.validateCode(code);
  if (!record) throw new Error('Invalid code');

  // 2. send request to add user to group/task
  // No supported apis so far

  // 3. return with request result;

  ctx.response.body = record;
  ctx.response.status = 200;
});

export default router;
