import React from  'react';
import { App, View, Navbar } from 'framework7-react';
import HomePage from './pages/home.jsx';
import AboutPage from './pages/about.jsx';

const routes = [{
  path: '/',
  component: HomePage,
}, {
  path: '/about',
  component: AboutPage,
}];

const rootPath = window.location.pathname.replace(/\/+$/, '');

export default () => (
  <App name="Credit" theme="auto" id="treeples.credit" routes={routes}>
    <Navbar title="Credit"></Navbar>
    <View
      main
      url="/"
      browserHistory
      browserHistorySeparator=""
      browserHistoryRoot={rootPath}
      animate={false}
    />
  </App>
);
