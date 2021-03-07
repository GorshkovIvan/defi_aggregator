import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import ResultsTable from '../ResultsTable.js';


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('ResultsTable');
    elem.setAttribute('id', 'resultsTable');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('resultsTable');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'ResultsTable test, renders table 1 (User specified Table)', () => {
    const elem = document.getElementById('resultsTable');

    console.log('HTML ->', document.body.innerHTML);

    var results = ResultsTable()    
    //const TableRowElem = elem.querySelector('Table.Row');
    //expect(TableRowElem).toHaveTextContent("hello?")

    //This test is currently doing nothing....
});

