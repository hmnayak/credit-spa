import React from "react";
import { App, View, Navbar } from "framework7-react";
import routes from "./routes";
import { logoutClicked, getCurUser } from "./services/authsvc";
import "../css/navbar.css";

const rootPath = window.location.pathname;

export default class Container extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isLoading: false,
      isAuthScreen: false,
      user: "",
    };
  }

  async componentDidMount() {
    await getCurUser().then((user) => {
      this.setState({ user: user });
    });
  }

  setAuthScreenLoaded(isAuthScreen) {
    this.setState({ isAuthScreen: isAuthScreen });
  }

  setLoading(isLoading) {
    this.setState({ isLoading: isLoading });
  }

  loading() {
    if (this.state.isLoading) {
      return <div className="center">Loading...</div>;
    }
  }

  credentialsContent() {
    if (this.state.user === "Guest") {
      return (
        <div className="right">
          <a href="/about/" className="link navlink">
            About
          </a>
          <a href="/login/" className="link navlink">
            Login
          </a>
          <a href="/signup/" className="link navlink">
            Signup
          </a>
        </div>
      );
    } else {
      return (
        <div className="right">
          <a href="/about/" className="link navlink">
            About
          </a>
          <a href="/signup/" className="link navlink">
            Signup
          </a>
          <a href="/" className="link navlink" onClick={logoutClicked}>
            Logout
          </a>
        </div>
      );
    }
  }

  userInfo() {
    return this.state.user;
  }

  headerContent() {
    if (this.state.isAuthScreen) {
      return <div className="page no-navbar no-toolbar" />;
    } else {
      return (
        <div className="navbar-inner">
          <a href="/" className="link">
            Credit
          </a>
          {this.loading()}
          {this.credentialsContent()}
        </div>
      );
    }
  }

  render() {
    return (
      <App
        name="Credit"
        theme="auto"
        id="treeples.credit"
        routes={routes(
          this.setLoading.bind(this),
          this.setAuthScreenLoaded.bind(this),
          this.headerContent.bind(this),
          this.userInfo.bind(this)
        )}
      >
        <Navbar>{this.headerContent()}</Navbar>
        <View
          main
          url={rootPath}
          browserHistory
          browserHistorySeparator=""
          pushState
          browserHistoryRoot=""
          animate={false}
          browserHistoryInitialMatch={false}
          browserHistoryStoreHistory={false}
        ></View>
      </App>
    );
  }
}
