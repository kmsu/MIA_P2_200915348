import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

import Navbar from './Components/navbar';
import VentanaComando from './Components/ventanaComandos';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Navbar></Navbar>
        <VentanaComando></VentanaComando>
      </div>
    );
  }
}

export default App;
