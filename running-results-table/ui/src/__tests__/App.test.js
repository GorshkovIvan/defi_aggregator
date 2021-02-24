import React from 'react';
import ReactDOM from 'react-dom';
import App from '..';

it('renders without crashing', () => {
    const div = document.createElement('div'); // create the div here
    ReactDOM.render(<App />, document.getElementById('root'));
});