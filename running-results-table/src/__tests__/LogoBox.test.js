import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import LogoBox from '../LogoBox.js';


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('LogoBox');
    elem.setAttribute('id', 'logoBox');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('logoBox');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'LogoBox renders as expected', () => {
    const elem = document.getElementById('logoBox');
    act( () => {
        render(<LogoBox/>, elem);
    });
});

