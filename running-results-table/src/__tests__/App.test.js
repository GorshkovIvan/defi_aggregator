import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import App from '../App';


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('app');
    elem.setAttribute('id', 'app');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('app');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'Full app renders without crashing', () => {
    const elem = document.getElementById('app');
    act( () => {
        render(<App/>, elem);
    });
    //console.log('HTML ->', document.body.innerHTML);
    //const h3Elem = elem.querySelector('h3');
    //expect(h3Elem).toHaveTextContent('Best cryptocurrencies to provide liquidity on are:')
});