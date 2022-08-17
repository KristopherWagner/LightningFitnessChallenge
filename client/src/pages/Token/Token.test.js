import { render } from '@testing-library/react';
import { rest } from 'msw';
import { setupServer } from 'msw/node';

import Token from './Token';

jest.mock('react-router-dom', () => ({
  useSearchParams: () => ([{
    get: () => '',
  }])
}));

const server = setupServer(
  rest.post('https://www.strava.com/oauth/token', (req, res, ctx) => {
    return res(ctx.json({ data: '' }));
  }),
);

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
  render(<Token />);
});

test('handles getting a token', async () => {
  render(<Token />);
});
