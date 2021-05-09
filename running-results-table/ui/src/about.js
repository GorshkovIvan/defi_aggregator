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
                    While in CeFi, infrastructure for benchmarking returns from various investments is well developed, in DeFi it is unregulated, 
                    and few marketplaces exist for comparing returns from available investments. Moreover, DeFi products work differently from their 
                    classic equivalents, and require alternative ways for calculating return on investment (ROI) and for quantifying risks. The 
                    engineering part of the problem is to develop a platform which aggregates and displays data on various pools, ranking them by ROI. 
                    The analytical part of the problem is to develop a standardised ROI formula, which will process data from pools into a single, 
                    reliable metric allowing investors to optimise their portfolio.
                </p>
            </div>
            
            <div class="Links">
                <div class="ui icon buttons">
                    <div class="ui button">
                        <a href="https://www.overleaf.com/project/606f024e52270d0b2fd71896">
                            <i class="file alternate icon" />
                        </a>
                    </div>
                    <div class="ui button">
                        <a href="https://gitlab.doc.ic.ac.uk/g207004212/defi_aggregator">
                            <i class="gitlab icon" />
                        </a>
                    </div>
                </div>
            </div>

            <div class="back-about">
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