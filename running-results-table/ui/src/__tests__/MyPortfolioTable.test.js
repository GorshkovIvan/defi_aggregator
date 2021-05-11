import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import MyPortfolioTable from '../MyPortfolioTable.js';

beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('MyPortfolioTable');
    elem.setAttribute('id', 'myPortfolioTable');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('myPortfolioTable');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'MyPortfolioTable test, renders table 1 (User specified Table)', () => {
    const elem = document.getElementById('myPortfolioTable');

    console.log('HTML ->', document.body.innerHTML);

    act( () => {
        render(<MyPortfolioTable/>, elem);
    });

});