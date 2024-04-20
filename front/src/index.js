import React from 'react';
//import ReactDOM from 'react-dom';
import { createRoot } from 'react-dom/client';
import App from './App';
import './index.css';
import 'bootstrap/dist/css/bootstrap.min.css';

/*ReactDOM.render(
  <App />,
  document.getElementById('root')
);*/

const root = createRoot(document.getElementById('root'));
root.render(<App />);