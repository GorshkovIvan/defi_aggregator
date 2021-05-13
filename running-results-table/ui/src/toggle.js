import React, { useEffect, useState } from 'react';
import { setTheme } from './themes';
import { Button } from 'semantic-ui-react';
import { Link } from 'react-router-dom';
import './App.css';
import logo from './Screenshot 2021-05-12 at 18.56.56.png'
import logo2 from './Screenshot 2021-05-12 at 18.56.39.png'


function Toggle() {
    const [togClass, setTogClass] = useState('dark');
    let theme = localStorage.getItem('theme');

    const handleOnClick = () => {
        if (localStorage.getItem('theme') === 'theme-dark') {
            setTheme('theme-light');
            setTogClass('light')
        } else {
            setTheme('theme-dark');
            setTogClass('dark')
        }
    }

    useEffect(() => {
        if (localStorage.getItem('theme') === 'theme-dark') {
            setTogClass('dark')
        } else if (localStorage.getItem('theme') === 'theme-light') {
            setTogClass('light')
        }
    }, [theme])

    return (
        <div className="ui container">
            <div className="container--toggle">
            <div className="logobox"><img src={localStorage.getItem('theme') === 'theme-dark' ? logo : logo2} width="25%"></img></div>
                <h1>High Yield 4 Me</h1>
                <h5>The place to find optimal yield for your cryptocurrency portfolio.*</h5>
                <div class="menuBar">
                    <Link to="/about" class="about">About</Link>
                    <Link to="/contact" class="contact">Contact</Link>
                </div>
                {
                    togClass === "light" ?
                    <Button id="toggle" className="toggle--checkbox" onClick={handleOnClick}>Dark Mode</Button>
                    :
                    <Button id="toggle" className="toggle--checkbox" onClick={handleOnClick}>Light Mode</Button>
                }
                <label htmlFor="toggle" className="toggle--label">
                    <span className="toggle--label-background"></span>
                </label>
            </div>
        </div>
    )
}

export default Toggle;
