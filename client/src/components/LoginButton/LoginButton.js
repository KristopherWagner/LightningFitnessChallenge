import Button from '@mui/material/Button';

export default function LoginButton() {
  const client_id = 45959;
  const redirect_uri = 'http://localhost:3000/';
  const url = `http://www.strava.com/oauth/authorize?client_id=${client_id}&response_type=code&redirect_uri=${redirect_uri}exchange_token&approval_prompt=force&scope=read,activity:read`;

  return (
    <Button href={url} variant='contained'>
      Authorize
    </Button>
  );
}