import React from 'react';
import { Form, Header, Segment, Button } from 'semantic-ui-react'
import styled from 'styled-components';

const sliderThumbStyles = (props) => (`
  width: 25px;
  height: 25px;
  background: ${props.color};
  cursor: pointer;
  outline: 5px solid #333;
  opacity: ${props.opacity};
  -webkit-transition: .2s;
  transition: opacity .2s;
`);

const Styles = styled.div`
  display: flex;
  align-items: center;
  color: #888;
  margin-top: 2rem;
  margin-bottom: 2rem;
  .value {
    flex: 1;
    font-size: 2rem;
  }
  .slider {
    flex: 6;
    -webkit-appearance: none;
    width: 100%;
    height: 15px;
    border-radius: 5px;
    background: #efefef;
    outline: none;
    &::-webkit-slider-thumb {
      -webkit-appearance: none;
      appearance: none;
      ${props => sliderThumbStyles(props)}
    }
    &::-moz-range-thumb {
      ${props => sliderThumbStyles(props)}
    }
  }
`;


export default class Slider2 extends React.Component {
    state = {
        risk_setting: ''
    };
        
  handleOnChange = (e) => this.setState({ risk_setting: e.target.value });
 
    onSubmit = this._onSubmit.bind(this);

    render() {
        return (
            <div className="ui container">
              <Form onSubmit={this.onSubmit}>
              
                  <Styles opacity={this.state.risk_setting > 10 ? (this.state.risk_setting / 10) : .1} color={this.props.color}>
                    <input type="range" min={0} max={10} value={this.state.risk_setting} className="slider" onChange={this.handleOnChange}/>                  
                  </Styles>                
                
                  <Button type='submit'>Submit</Button>

              </Form>
            </div>
        );
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