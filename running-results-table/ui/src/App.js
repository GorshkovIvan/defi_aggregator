import logo from './eth.png';

import React, { Component } from 'react';
import ConnectedResultsTable from './ConnectedResultsTable';
import NewResultsForm from './NewResultsForm';
import RankedCurrenciesTable from './RankedCurrenciesTable';
import './App.css';

class App extends Component {

  render() {

    return (

    <div className="App">
    <div className ="App2"><h1> Welcome to HighYield4Me! </h1> </div>

    <div class="center"><img src={logo} alt="our logo" className="Applogo" class="center" width="1%" /> </div>

        <ConnectedResultsTable />
        <NewResultsForm />
        <div className ="App3"> </div>
        <RankedCurrenciesTable />
      </div>
    );
  }
}
export default App;
