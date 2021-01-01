import React, { useState } from "react";
import { App, View, Navbar, f7 } from "framework7-react";
import { user, setNavigate } from "./authsvc";
import routes from "./routes";

const rootPath = window.location.pathname.replace(/\/+$/, "");

export default class Container extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      isLoading: false,
    };
  }

  componentDidMount() {
    setNavigate(f7.views.main.router.navigate.bind(f7.views.main.router));
  }
  
  setLoading(isLoading) {
    this.setState({ isLoading: isLoading });
  }

  loading() {
    if (this.state.isLoading) {
      return <div>Loading...</div>;
    }
  }

  render() {
    return (
      <App
        name="Credit"
        theme="auto"
        id="treeples.credit"
        routes={routes(this.setLoading.bind(this))}
      >
        <Navbar title="Credit">{this.loading()}</Navbar>
        <View
          main
          url={rootPath}
          browserHistory
          browserHistorySeparator=""
          browserHistoryRoot=""
          animate={false}
        />
      </App>
    );
  }

}
