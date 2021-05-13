import React from "react";
import { Link } from 'react-router-dom';
import { useEffect } from 'react';
import { keepTheme } from './themes.js';

function About() {
    
    useEffect(() => { keepTheme(); })

    return (
        <div class="About">
            <h1 className="about-hyfm">High Yield 4 Me</h1>
            <div className="about-header">About High Yield 4 Me</div>
            <div class="paragraph">
                <p>Decentralised finance offers a new way to manage your finances: without banks or other intermediaries.</p>
                <p>The function of banks is performed by 'Liquidity Pools' where you can deposit your cryptocurrencies and earn a yield, similar to a traditional bank deposit.</p>
                <p>But how should you choose which pools to invest in? Many factors must be considered: how secure the pool is, how volatile the returns are, and how stable the tokens are.
                    These metrics are less readily available currently in decentralised finance as opposed to traditional finance, and our objective with this project is to make the process of choosing Liquidity Pools easier, more transparent, and help users to choose the best pool for their deposits.</p>
                <p>Feel free to browse our report and GitLab repository for more information!</p>
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