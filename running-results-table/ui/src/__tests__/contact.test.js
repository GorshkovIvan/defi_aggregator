import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import Contact from '../contact';


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('Contact');
    elem.setAttribute('id', 'contact');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('contact');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'Contact page renders without crashing', () => {
    const elem = document.getElementById('contact');
    act( () => {
        render(<Contact/>, elem);
    });
    //console.log('HTML ->', document.body.innerHTML);
    //const h3Elem = elem.querySelector('h3');
    //expect(h3Elem).toHaveTextContent('Best cryptocurrencies to provide liquidity on are:')
});