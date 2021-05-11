import React, { useEffect, useState } from 'react';
import { setTheme } from './themes';
import { Button } from 'semantic-ui-react';
import './App.css';

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
