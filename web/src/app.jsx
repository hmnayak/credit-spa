import React from  'react';
import { App, View, Navbar } from 'framework7-react';
import HomePage from './pages/home.jsx';
import AboutPage from './pages/about.jsx';

const params = {
  name: 'Credit',
  theme: 'auto',
  id: 'treeples.credit',
  routes: [{
    path: '/',
    component: HomePage,
  },{
    path: '/about',
    component: AboutPage,
  }],  
};

const rootPath = window.location.pathname.replace(/\/+$/, '');

export default () => (
  <App params={params}>
    <Navbar title="Credit"></Navbar>
    <View main pushState={true} pushStateSeparator="" pushStateRoot={rootPath} />
  </App>
);
