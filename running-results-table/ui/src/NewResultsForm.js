import React from 'react';
import { Form, Header, Segment, Button } from 'semantic-ui-react'

// does not do any validation
export default class NewResultsForm extends React.Component {
    state = {
        name: '',
        time: '',
        pool_sz: ''
    };
    onChangeName = this._onChangeName.bind(this);
    onChangeTime = this._onChangeTime.bind(this);
    onChangePool_sz = this._onChangePool_sz.bind(this);
    onSubmit = this._onSubmit.bind(this);
    render() {
        return (
            <div className="ui container">
                <Segment vertical>
                    <Header>New Result</Header>
                    <Form onSubmit={this.onSubmit}>
                        <Form.Field>
                            <label>Name</label>
                            <input placeholder='Name' value={this.state.name} onChange={this.onChangeName} />
                        </Form.Field>
                        <Form.Field>
                            <label>Time</label>
                            <input placeholder='Time' value={this.state.time} onChange={this.onChangeTime} />
                        </Form.Field>
                        <Form.Field>
                            <label>Pool Size</label>
                            <input placeholder='Pool Size' value={this.state.pool_sz} onChange={this.onChangePool_sz} />
                        </Form.Field>
                        <Button type='submit'>Submit</Button>
                    </Form>
                </Segment>
            </div>
        );
    }
    _onChangeName(e) {
        this.setState({
            name: e.target.value
        });
    }
    _onChangeTime(e) {
        this.setState({
            time: e.target.value
        });
    }
    _onChangePool_sz(e) {
        this.setState({
            pool_sz: e.target.value
        });
    }
    _onSubmit() {
        const payload = {
            name: this.state.name,
            time: parseFloat(this.state.time),
            pool_sz: parseFloat(this.state.pool_sz)
        };
        fetch('http://localhost:8080/results', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
        });
        this.setState({
            name: '',
            time: '',
            pool_sz: ''
        });
    }
}