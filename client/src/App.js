import logo from './logo.svg';
import './App.css';
import { LoginButton } from './components';

export default function App() {
    return (
        <div className="App">
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" />
                <p>
          This page is a work in progress
                </p>
                <LoginButton />
            </header>
        </div>
    );
}
