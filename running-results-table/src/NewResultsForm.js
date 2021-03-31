import React from 'react';
import { Form, Header, Segment, Button } from 'semantic-ui-react'


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
                <Segment vertical>
                    <Header>Please Enter Your Cryptocurrency Portfolio and we will suggest how to maximize its yield:</Header>
                    <Form onSubmit={this.onSubmit}>

                        <div class="mainDIV">
                            <div class="leftDIV">
                                <Form.Field>
                                    <label>Token</label>
                                    <input placeholder='ETH' value={this.state.token} onChange={this.onChangeName} />
                                </Form.Field>
                            </div>
                            <div class="middleDIV">
                                <Form.Field>
                                    <label>Amount</label>
                                    <input placeholder='Amount' value={this.state.amount} onChange={this.onChangeTime} />
                                </Form.Field>
                            </div>
                            <div class="rightDIV">
                                <Form.Field>
                                    <label> Add </label>
                                    <Button type='submit'>Submit</Button>
                                </Form.Field>
                            </div> 
                        </div>
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