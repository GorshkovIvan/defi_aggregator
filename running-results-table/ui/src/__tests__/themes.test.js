import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import themes from '../themes';


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('Themes');
    elem.setAttribute('id', 'themes');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('themes');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'Full themes renders without crashing', () => {
    const elem = document.getElementById('themes');
    act( () => {
        render(<themes/>, elem);
    });
    //console.log('HTML ->', document.body.innerHTML);
    //const h3Elem = elem.querySelector('h3');
    //expect(h3Elem).toHaveTextContent('Best cryptocurrencies to provide liquidity on are:')
});