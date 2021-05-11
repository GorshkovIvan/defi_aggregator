import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import Toggle from '../toggle';


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('Toggle');
    elem.setAttribute('id', 'toggle');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('toggle');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'Toggle renders without crashing', () => {
    const elem = document.getElementById('toggle');
    act( () => {
        render(<Toggle/>, elem);
    });
    //console.log('HTML ->', document.body.innerHTML);
    //const h3Elem = elem.querySelector('h3');
    //expect(h3Elem).toHaveTextContent('Best cryptocurrencies to provide liquidity on are:')
});