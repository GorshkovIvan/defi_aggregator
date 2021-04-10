import React from "react";
import { Link } from 'react-router-dom';
import { useEffect } from 'react';
import { keepTheme } from './themes.js';

function Contact() {

    useEffect(() => { keepTheme(); })

    return (
        <div class="Contact">
            <h1>High Yield 4 Me</h1>
            <h4>Contact Us</h4>
            <div class="back">
                <Link to="/">
                    <div class="ui animated fade button" tabindex="0">
                        <div class="visible content">Back</div>
                        <div class="hidden content">
                            <i class="left arrow icon" />
                        </div>
                    </div>
                </Link>
            </div>
        </div>
    );
}
export default Contact;