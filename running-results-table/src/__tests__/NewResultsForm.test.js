import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import NewResultForm from '../NewResultsForm.js';


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('NewResultForm');
    elem.setAttribute('id', 'newResultForm');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('newResultForm');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'NewResultForm test, fetch result from user input on webpage', () => {
    const elem = document.getElementById('newResultForm');
    act( () => {
        render(<NewResultForm/>, elem);
    });
});