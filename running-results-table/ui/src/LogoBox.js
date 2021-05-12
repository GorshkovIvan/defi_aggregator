import React from 'react';
import ReactDOM from 'react-dom';
//import logo from './Screenshot 2021-05-12 at 18.56.56.png'
import logo from './Screenshot 2021-05-12 at 18.56.39.png'


class LogoBox extends React.Component {
    render() {
      return <div className="logobox"><img src={logo} width="10%"></img></div>;
    }
  }

  ReactDOM.render(<LogoBox />, document.getElementById('root') || document.createElement('div'));
  export default LogoBox;