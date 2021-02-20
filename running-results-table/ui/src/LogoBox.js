import React from 'react';
import ReactDOM from 'react-dom';
import logo from './eth.png';

class LogoBox extends React.Component {
    render() {
      return <div className="Applogo"> <img src={logo} width="1%"/> </div>;
    }
  }
  
  ReactDOM.render(<LogoBox />, document.getElementById('root'));
  export default LogoBox;