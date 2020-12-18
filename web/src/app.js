import React from  'react';
import ReactDOM from  'react-dom';
import Framework7 from 'framework7';
import Framework7React from 'framework7-react';
import App from './app.jsx';

Framework7.use(Framework7React);

ReactDOM.render(React.createElement(App), document.getElementById('app'));
