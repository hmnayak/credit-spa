import React from "react";
import { App, View, Navbar } from "framework7-react";
import routes from "./routes";

const rootPath = window.location.pathname.replace(/\/+$/, "");

export default class Container extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isLoading: false,
    };
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
          browserHistoryInitialMatch={false}
        />
      </App>
    );
  }
}
