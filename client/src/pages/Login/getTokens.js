import axios from 'axios';

export default async function getTokens({ code }) {
  const response = await axios.post('http://localhost:8080/token', { code }, { headers: {
    'Access-Control-Allow-Origin': 'http://localhost:3000',
  }});

  return response.data;
}