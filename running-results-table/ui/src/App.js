import React, { useEffect } from 'react';
import ConnectedResultsTable from './ConnectedResultsTable';
import NewResultsForm from './NewResultsForm';
import RankedCurrenciesTable from './RankedCurrenciesTable';
import './App.css';
import './theme.css';
import Slider from './Slider.js';
import Toggle from './toggle.js';
import { keepTheme } from './themes.js';
import './about.js';
import './contact.js';
import ConnectedMyPortfolioTable from './ConnectedMyPortfolioTable';

function App() {
  
  useEffect(() => { keepTheme(); })

  return (
    
    // ======================= !!! IMPORTANT !!! ============================
    // commented out for the website to work, uncomment for the tests to work

    //<Router>
    <div className="App">
      <div className="TopTable">
        <Toggle className="logobox"></Toggle>
      </div>
      <RankedCurrenciesTable />

      <div className="MiddleDivider"></div>

      <div class="ui container">
        <div class="floatContainer">
          <div class="recommendedPortfolio">
            <div class="portfolio">
              <ConnectedMyPortfolioTable />
            </div>
            <div class="portfolio">
              <ConnectedResultsTable />
            </div>
          </div>
          <div class="resultsForm">
            <NewResultsForm />
            <Slider />
          </div>
        </div>
      </div>
      <div><h6>*Disclaimer: the information presented on this website is not financial advice. Invest at your own risk. High Yield 4 Me is not responsible for any losses incurred from using this website.</h6></div>
    </div>
    
    //</Router>
  );
}

export default App;
