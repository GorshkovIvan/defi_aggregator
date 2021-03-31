import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import RankedCurrenciesTable from '../RankedCurrenciesTable.js';


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('RankedCurrenciesTable');
    elem.setAttribute('id', 'rankedCurrenciesTable');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('rankedCurrenciesTable');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'RankedCurrenciesTable test, fetches data from backend for ResultsTable2', () => {
    const elem = document.getElementById('rankedCurrenciesTable');
    act( () => {
        render(<RankedCurrenciesTable/>, elem);
    });
});