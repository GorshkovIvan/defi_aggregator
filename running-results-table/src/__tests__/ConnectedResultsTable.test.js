import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import ConnectedResultsTable from '../ConnectedResultsTable';


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('ConnectedResultsTable');
    elem.setAttribute('id', 'connectedResultsTable');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('connectedResultsTable');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'ConnectedResultsTable test, links front-end to back-end', () => {
    const elem = document.getElementById('connectedResultsTable');
    act( () => {
        render(<ConnectedResultsTable/>, elem);
    });
});