import { act, render } from '@testing-library/react';
import { rest } from 'msw';
import { setupServer } from 'msw/node';

import Login from './Login';

jest.mock('react-router-dom', () => ({
  useSearchParams: () => ([{
    get: () => '',
  }])
}));

const server = setupServer();
beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

test('handles server error', async () => {
  server.use(
    rest.post('https://www.strava.com/oauth/token', (req, res, ctx) => (
      res(
        ctx.status(401),
        ctx.json({ errorMessage: 'error' })
      )
    )),
  );
  await act(async () => {
    await render(<Login />);
  });
});

test('handles getting a token', async () => {
  server.use(
    rest.post('https://www.strava.com/oauth/token', (req, res, ctx) => (
      res(ctx.status(200), ctx.json({ athlete: {} }))
    )),
  );
  await act(async () => {
    await render(<Login />);
  });
});
