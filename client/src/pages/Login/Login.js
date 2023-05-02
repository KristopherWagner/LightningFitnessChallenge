import { useEffect, useState } from 'react';
import { useSearchParams } from 'react-router-dom';

import Alert from '@mui/material/Alert';
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardHeader from '@mui/material/CardHeader';
import CircularProgress from '@mui/material/CircularProgress';
import Snackbar from '@mui/material/Snackbar';
import Typography from '@mui/material/Typography';

import getTokens from './getTokens';

export default function Login() {
  const [params,] = useSearchParams();
  const code = params.get('code');

  const [token, setToken] = useState(null);
  const [isSnackbarOpen, setIsSnackbarOpen] = useState(false);
  const closeSnackbar = () => setIsSnackbarOpen(false);

  useEffect(() => {
    const loadTokens = async () => {
      try {
        const response = await getTokens({ code });
        setToken(response);
      } catch (error) {
        setIsSnackbarOpen(true);
      }
    };

    if (typeof (code) === 'string' && code.length > 0) {
      loadTokens();
    }
  }, [code]);

  return (
    <Box sx={{
      display: 'flex',
      justifyContent: 'space-around',
      margin: '24px auto',
      maxWidth: '80vw',
    }}>
      {token == null ? (
        <CircularProgress />
      ) : (
        <Card>
          <CardHeader title='Your token' />
          <CardContent>
            <Typography role="info">
              {`Access: ${token.access_token}`}
            </Typography>
            <Typography role="info">
              {`Expires: ${token.expires_at}`}
            </Typography>
            <Typography role="info">
              {`ID: ${token.athlete.id}`}
            </Typography>
            <Typography role="info">
              {`Refresh: ${token.refresh_token}`}
            </Typography>
          </CardContent>
        </Card>
      )}
      <Snackbar
        autoHideDuration={6000}
        anchorOrigin={{ horizontal: 'center', vertical: 'bottom' }}
        open={isSnackbarOpen}
        onClose={closeSnackbar}
      >
        <Alert onClose={closeSnackbar} severity='error'>Unable to get token</Alert>
      </Snackbar>
    </Box>
  );
}
