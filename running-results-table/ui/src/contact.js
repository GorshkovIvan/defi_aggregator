import React from "react";
import { Link } from 'react-router-dom';
import { useEffect } from 'react';
import { keepTheme } from './themes.js';

function Contact() {

    useEffect(() => { keepTheme(); })

    return (
        <div class="Contact">
            <h1>High Yield 4 Me</h1>
            <h4>Meet the Team:</h4>
            <h3>Some kind of description...</h3>
            <div class="ui container">
                <h4>Databases & Testing</h4>
                <div class="row">
                    <div class="column">
                        <h3>Aaron Lam</h3>
                        <a href="mailto:ayl2617@ic.ac.uk">Send me spam.</a>
                    </div>
                    <div class="column">
                        <h3>Marcus Loo</h3>
                        <a href="mailto:mzl20@ic.ac.uk">Send me spam.</a>
                    </div>
                </div>

                <h4>Front End</h4>
                <div class="row">
                    <div class="column">
                        <h3>Chris Stanford</h3>
                        <a href="mailto:cas220@ic.as.uk">Send me spam.</a>
                    </div>
                    <div class="column">
                        <h3>John Kwon</h3>
                        <a href="mailto:jk3120@ic.ac.uk">Send me spam.</a>
                    </div>
                </div>

                <h4>Back End</h4>
                <div class="row">
                    <div class="column">
                        <h3>Ivan Gorshkov</h3>
                        <a href="mailto:ig420@imperial.ac.uk">Send me spam.</a>
                    </div>
                    <div class="column">
                        <h3>Yaroslav Afonin</h3>
                        <a href="mailto:ya1220@imperial.ac.uk">Send me spam.</a>
                    </div>
                </div>
            </div>

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