import axios from 'axios';

export default async function getTokens({ code }) {
  const response = await axios.post('https://www.strava.com/oauth/token', {
    client_id: 45959,
    client_secret: '',
    code,
    grant_type: 'authorization_code',
  });

  return response.data;
}