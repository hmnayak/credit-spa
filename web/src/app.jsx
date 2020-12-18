import React from  'react';
import { App, View, Page, Navbar } from 'framework7-react';

export default () => (
  <App params={{ theme: 'auto', name: 'Credit', id: 'treeples.credit' }}>

    <View main>

      <Page>
        <Navbar title="Credit"></Navbar>
        <p>Hello world</p>
      </Page>

    </View>
  </App>
);