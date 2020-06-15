import * as Koa from 'koa';

const app = new Koa();

app.use(
  async (ctx: Koa.ParameterizedContext<Koa.DefaultState, Koa.DefaultContext>): Promise<void> => {
    ctx.body = 'Hello World';
  },
);

app.listen(3000);
