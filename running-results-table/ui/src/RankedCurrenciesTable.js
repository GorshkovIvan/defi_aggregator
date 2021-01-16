import React from 'react';
import ResultsTable2 from './ResultsTable2';
import Pusher from 'pusher-js';

const socket = new Pusher('7885860875bb513c3e34', {
    cluster: 'eu',
    encrypted: true
});

export default class ConnectedResultsTable2 extends React.Component {
    state = {
        results: []
    };
    componentDidMount() {
        const channel = socket.subscribe('results');
        channel.bind('results', (data) => {
            this.setState(data);
        });

        // change this url:

        fetch('http://localhost:8080/results')
            .then((response) => response.json())
            .then((response) => this.setState(response));
    }
    render() {
        return <ResultsTable2 results={this.state.results} />;
    }
}