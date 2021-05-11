import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import About from '../about';


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('About');
    elem.setAttribute('id', 'about');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('about');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'About page renders without crashing', () => {
    const elem = document.getElementById('about');
    act( () => {
        render(<About/>, elem);
    });
    //console.log('HTML ->', document.body.innerHTML);
    //const h3Elem = elem.querySelector('h3');
    //expect(h3Elem).toHaveTextContent('Best cryptocurrencies to provide liquidity on are:')
});