import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import TextBox from '../TextBox.js';


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('TextBox');
    elem.setAttribute('id', 'textBox');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('textBox');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'TextBox renders as expected', () => {
    const elem = document.getElementById('textBox');
    act( () => {
        render(<TextBox/>, elem);
    });
});