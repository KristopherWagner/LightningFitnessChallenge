import { render } from '@testing-library/react';
import LoginButton from './LoginButton';

test('renders without crashin', () => {
  render(<LoginButton />);
});