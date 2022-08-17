import { useEffect, useState } from 'react';
import axios from 'axios';
import { useSearchParams } from 'react-router-dom';

async function getTokens({ code }) {
  const response = await axios.post('https://www.strava.com/oauth/token', {
    client_id: 45959,
    client_secret: '',
    code,
    grant_type: 'authorization_code',
  });

  return response.data;
}

export default function Token() {
  const [params, ] = useSearchParams();
  const [token, setToken] = useState(null);
  const code = params.get('code');

  useEffect(() => {
    const loadTokens = async () => {
      try {
        const response = await getTokens({ code });
        setToken(response);
      } catch(error) {
        console.error('Unable to get token');
      }
    };
    loadTokens();
  }, [code]);
  return <p role="info">{JSON.stringify(token, 0, 2)}</p>;
}