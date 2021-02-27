import React from 'react';
import ReactDOM from 'react-dom';

class TextBox extends React.Component {
    render() {
      return <div className="HeaderBox"> Welcome to High Yield 4 Me! </div>;
    }
  }
  
  ReactDOM.render(<TextBox />, document.getElementById('root') || document.createElement('div'))
  export default TextBox;