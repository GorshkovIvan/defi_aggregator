import { useEffect } from 'react';
import ConnectedResultsTable from './ConnectedResultsTable';
import NewResultsForm from './NewResultsForm';
import RankedCurrenciesTable from './RankedCurrenciesTable';
import './App.css';
import './theme.css';
import Slider from './Slider.js';
import Toggle from './toggle.js';
import { keepTheme } from './themes.js';

function App() {
  
  useEffect(() => { keepTheme(); })

  return (

    <div className="App">
      <div className="TopTable">
        <h1>High Yield 4 Me</h1>
        <h5>The place to find optimal yield for your cryptocurrency portfolio.*</h5>
      </div>
      <Toggle></Toggle>
      <RankedCurrenciesTable />

      <div className="MiddleDivider"></div>

      <div class="ui container">
        <div class="floatContainer">
          <div class="recommendedPortfolio">
            <ConnectedResultsTable />
          </div>
          <div class="resultsForm">
            <NewResultsForm />
            <Slider />
          </div>
        </div>
      </div>
      <div><h6>*Disclaimer: the information presented on this website is not financial advice. Invest at your own risk. High Yield 4 Me is not responsible for any losses incurred from using this website.</h6></div>
    </div>
  );
}

export default App;
