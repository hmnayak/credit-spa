import React from "react";
import { App, View, Navbar, Link } from "framework7-react";
import routes from "./routes";
import { logoutClicked, getCurUser } from "./services/authsvc";

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
      return <div class="center">Loading...</div>;
    }
  }
  
  credentialsContent(){
    if(getCurUser() === "Guest"){
      return (
        <div class="right">
            <a href="/login/" class="link">Login</a>
            <a href="/signup/" class="link">Signup</a>
        </div>
      );
    }else {
      return(
        <div class="right">
          <a href="/" class="link" onClick={logoutClicked}>Logout</a>
      </div>
      )
    }
  }
  
  headerContent() {
      return (
        <div class="navbar-inner">
          <div class="title">Credit</div>
          {this.loading()}
          {this.credentialsContent()}
      </div>);
  }
  
  render() {
    return (
      <App
        name="Credit"
        theme="auto"
        id="treeples.credit"
        routes={routes(this.setLoading.bind(this))}
      >
        <Navbar>{this.headerContent()}</Navbar>
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
