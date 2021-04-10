import React from "react";
import { Link } from 'react-router-dom';
import { useEffect } from 'react';
import { keepTheme } from './themes.js';

function About() {
    
    useEffect(() => { keepTheme(); })

    return (
        <div class="About">
            <h1>High Yield 4 Me</h1>
            <h4>About High Yield 4 Me</h4>
            <div class="paragraph">
                <p>
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec aliquet sapien orci, vitae vestibulum ligula malesuada commodo. Suspendisse dapibus aliquam justo. Sed enim metus, feugiat vitae sollicitudin vitae, finibus at tortor. Interdum et malesuada fames ac ante ipsum primis in faucibus. Maecenas accumsan eleifend malesuada. Fusce felis arcu, elementum eget auctor id, fringilla a nisi. Integer in arcu purus. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Maecenas malesuada dui et est rutrum, sit amet laoreet elit venenatis. Praesent ultrices tincidunt pretium. Nam vestibulum, dolor in ultricies iaculis, mi lorem dapibus ligula, laoreet consequat augue odio at mauris. Donec a laoreet leo. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.
                </p>
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
export default About;