import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import ConnectedMyPortfolioTable from "../ConnectedMyPortfolioTable";


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('ConnectedMyPortfolioTable');
    elem.setAttribute('id', 'connectedMyPortfolioTable');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('connectedMyPortfolioTable');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'ConnectedMyPortfolioTable test, links front-end to back-end', () => {
    const elem = document.getElementById('connectedMyPortfolioTable');
    act( () => {
        render(<ConnectedMyPortfolioTable/>, elem);
    });
});