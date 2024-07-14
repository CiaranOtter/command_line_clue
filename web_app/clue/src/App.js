import logo from './logo.svg';
import './App.css';

import { loadSync } from '@grpc/proto-loader'
import { loadPackageDefinition } from '@grpc/grpc-js'

const packageDef = loadSync("/Users/ciaranotter/Documents/personal/command_line_clue/proto/clue.proto")
const requestProto = loadPackageDefinition(packageDef)

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
