import React from 'react';
import { Form, Header, Segment, Button } from 'semantic-ui-react'

export default class Slider2 extends React.Component {
    state = {
        risk_setting: ''
    };
    onChangeTime = this._onChangeTime.bind(this);
    
    onSubmit = this._onSubmit.bind(this);
    render() {
        return (
            <div className="ui container">
                <Segment vertical>
                    <Form onSubmit={this.onSubmit}>
                        <Form.Field>
                            <input placeholder='Risk setting' value={this.state.risk_setting} onChange={this.onChangeTime} />
                        </Form.Field>

                        <Button type='submit'>Submit</Button>
                    </Form>
                </Segment>
            </div>
        );
    }

    _onChangeTime(e) {
        this.setState({
          risk_setting: e.target.value
        });
    }

    _onSubmit() {
        const payload = {
            risk_setting: parseFloat(this.state.risk_setting)
        };
        fetch('http://localhost:8080/results2', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
        });
        this.setState({
          risk_setting: ''
        });
    }
}