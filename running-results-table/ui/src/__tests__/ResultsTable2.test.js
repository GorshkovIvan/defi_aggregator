import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import ResultsTable2 from '../ResultsTable.js';
import ResultsTable from "../ResultsTable.js";


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('ResultsTable2');
    elem.setAttribute('id', 'resultsTable2');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('resultsTable2');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'ResultsTable2 test, renders table 2 (Ranked currency pairs table)', () => {
    const elem = document.getElementById('resultsTable2');
    console.log('HTML ->', document.body.innerHTML);
    act( () => {
        render(<ResultsTable2/>, elem);
    });
});
