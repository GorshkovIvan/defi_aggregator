import React from 'react';
import { Form, Segment, Button, Popup } from 'semantic-ui-react'

// does not do any validation
export default class NewResultsForm extends React.Component {
    state = {
        token: '',
        amount: '',
//        pool_sz: ''
    };
    onChangeName = this._onChangeName.bind(this);
    onChangeTime = this._onChangeTime.bind(this);
  //  onChangePool_sz = this._onChangePool_sz.bind(this);
    onSubmit = this._onSubmit.bind(this);
    render() {
        return (
            <div className="ui container">
                <Segment>
                    <h3>Enter Your Portfolio <Popup content="Our engine will suggest the best-yielding allocation to available pools based on your current cryptocurrency portfolio." position="top center" trigger={<i class="info circle icon portfolio-popup"></i>} /></h3>
                    <Form onSubmit={this.onSubmit}>
                        <Form.Field>
                            <label><div className="inputLabel">Token</div></label>
                            <select value={this.state.token} onChange={this.onChangeName}>
                                <option value="" selected disabled hidden>Select Token</option>
                                <option value="ETH">ETH</option>
                                <option value="DAI">DAI</option>
                                <option value="UNI">UNI</option>
                            </select>
                        </Form.Field>
                        <Form.Field>
                            <label><div className="inputLabel">Amount</div></label>
                            <input type="number" step="any" placeholder='Enter Amount' value={this.state.amount} onChange={this.onChangeTime} />
                        </Form.Field>
                        <Form.Field className="portfolioSubmit">
                            <Button type='submit'>Submit</Button>
                        </Form.Field>
                    </Form>
                </Segment>
            </div>
        );
    }
    _onChangeName(e) {
        this.setState({
            token: e.target.value
        });
    }
    _onChangeTime(e) {
        this.setState({
            amount: e.target.value
        });
    }
/*    _onChangePool_sz(e) {
        this.setState({
            pool_sz: e.target.value
        });
    }*/
    _onSubmit() {
        const payload = {
            token: this.state.token,
            amount: parseFloat(this.state.amount),
           // pool_sz: parseFloat(this.state.pool_sz)
        };
        fetch('http://localhost:8080/results', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
        });
        this.setState({
            token: '',
            amount: '',
        //    pool_sz: ''
        });
    }
}

/*
                        <Form.Field>
                            <label>Pool Size</label>
                            <input placeholder='Pool Size' value={this.state.pool_sz} onChange={this.onChangePool_sz} />
                        </Form.Field>
*/