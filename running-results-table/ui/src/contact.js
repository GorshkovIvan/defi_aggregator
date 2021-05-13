import React from "react";
import { Link } from 'react-router-dom';
import { useEffect } from 'react';
import { keepTheme } from './themes.js';

function Contact() {

    useEffect(() => { keepTheme(); })

    return (
        <div class="Contact">
            <h1 className="about-hyfm">High Yield 4 Me</h1>
            <div className="about-header">Meet the Team:</div>
            <div class="paragraph">
                <p>We are a group of MSc Computer Science students at Imperial College London. We are cryptocurrency enthusiasts and have created this project to help cryptocurrency users make better choices for earning a yield from their investments.</p>
                <p>We are always eager to hear from potential investors who share the same vision as we do in revolutionising the cryptocurrency space. Please let us know by emailing one of our team members below!</p>
            </div>
            <div class="paragraph contact-container">
                <div class="row">
                    <div class="column-left">
                        <h3>Aaron Lam</h3>
                        <a href="mailto:ayl2617@ic.ac.uk">Email Aaron</a>
                    </div>
                    <div class="column-right">
                        <h3>Marcus Loo</h3>
                        <a href="mailto:mzl20@ic.ac.uk">Email Marcus</a>
                    </div>
                </div>

                <div class="row">
                    <div class="column-left">
                        <h3>Chris Stanford</h3>
                        <a href="mailto:cas220@ic.as.uk">Email Chris</a>
                    </div>
                    <div class="column-right">
                        <h3>John Kwon</h3>
                        <a href="mailto:jk3120@ic.ac.uk">Email John</a>
                    </div>
                </div>

                <div class="row">
                    <div class="column-left">
                        <h3>Ivan Gorshkov</h3>
                        <a href="mailto:ig420@imperial.ac.uk">Email Ivan</a>
                    </div>
                    <div class="column-right">
                        <h3>Yaroslav Afonin</h3>
                        <a href="mailto:ya1220@imperial.ac.uk">Email Yaroslav</a>
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