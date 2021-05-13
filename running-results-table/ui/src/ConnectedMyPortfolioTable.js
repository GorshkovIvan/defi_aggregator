import React from 'react';
import MyPortfolioTable from './MyPortfolioTable.js';
import Pusher from 'pusher-js';
//import 'semantic-ui-css/semantic.js';
const socket = new Pusher('7885860875bb513c3e34', {
    cluster: 'eu',
    encrypted: true
});

export default class ConnectedMyPortfolioTable extends React.Component {
    state = {
        results_original: []
    };
    componentDidMount() {
        const channel = socket.subscribe('results_original');
        channel.bind('results_original', (data2) => {
            this.setState(data2);
        });


        // change this url:

        fetch('http://localhost:8080/results')
            .then((response) => response.json())
            .then((response) => this.setState(response));
    }
    render() {
        return <MyPortfolioTable results={this.state.results_original} />;
    }
}
		