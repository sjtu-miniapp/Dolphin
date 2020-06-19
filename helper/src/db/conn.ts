import * as pg from 'pg-promise';

import * as config from '../../config.json';

export const sharecode: pg.IDatabase<any> = pg()(config.database);
