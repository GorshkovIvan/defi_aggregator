import React from 'react';
import {act} from 'react-dom/test-utils'
import {render, unmountComponentAtNode} from 'react-dom';
import Slider from '../Slider.js';


beforeEach( ()=> { //before each test create a div element
    const elem = document.createElement('Slider');
    elem.setAttribute('id', 'slider');
    document.body.appendChild(elem);
});

afterEach( ()=> { //remove div element so next test has clean <body>
    const elem = document.getElementById('slider');
    unmountComponentAtNode(elem);
    elem.remove();
})

test( 'Slider renders as expected', () => {
    const elem = document.getElementById('slider');
    act( () => {
        render(<Slider/>, elem);
    });
});
