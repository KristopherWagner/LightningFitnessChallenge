import './App.css';
import { Header, LoginButton } from './components';

import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardHeader from '@mui/material/CardHeader';

export default function App() {
  return (
    <div className="App">
      <Header />
      <Box sx={{ display: 'flex', justifyContent: 'space-around', padding: '24px' }}>
        <Card>
          <CardHeader title='Lightning Fitness Challnege' subheader='This page is a work in progress'/>
          <CardActions sx={{ justifyContent: 'space-around' }}>
            <LoginButton />
          </CardActions>
        </Card>
      </Box>
    </div>
  );
}
