import React from 'react';
import { render, screen } from '@testing-library/react';
import App from '..';

it('renders title', () => {
    render(<App />);
    expect(screen.getByText(' Welcome to High Yield 4 Me!')).toBeInTheDocument();
});