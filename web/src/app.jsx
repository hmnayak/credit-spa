import React from "react";
import { App, View, Navbar } from "framework7-react";
import routes from "./routes";
import { logoutClicked, getUsername, onUserChange } from "./services/authsvc";
import ErrorBoundary from "./pages/error";
import { NotificationMsg } from "./components/notification";
import "../css/navbar.css";

const rootPath = window.location.pathname;

export default class Container extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      isLoading: false,
      user: "",
      message: "",
    };
  }

  componentDidMount() {
    onUserChange(this.triggerUserChange.bind(this));
    this.triggerUserChange();
  }

  triggerLoading(isLoading) {
    this.setState({ isLoading: isLoading });
  }

  triggerNotification(message) {
    this.setState({ message: message });
  }

  triggerUserChange() {
    getUsername().then((username) => {
      this.setState({ username: username });
    });
  }

  onLogoutClick() {
    logoutClicked().then(() => {
      window.location.reload();
    });
  }

  renderLoading() {
    if (this.state.isLoading) {
      return <div className="center">Loading...</div>;
    }
  }

  renderSession() {
    if (this.state.username) {
      return (
        <>
          <span>{this.state.username}</span>
          <a
            href="/"
            className="link navlink"
            onClick={this.onLogoutClick.bind(this)}
          >
            Logout
          </a>
        </>
      );
    }
    return (
      <>
        <a href="/login/" className="link navlink">
          Login
        </a>
        <a href="/signup/" className="link navlink">
          Signup
        </a>
      </>
    );
  }

  renderNotification() {
    if (this.state.message) {
      return <NotificationMsg successMsg={this.state.message} />;
    }
  }

  render() {
    return (
      <ErrorBoundary>
        <App
          name="Credit"
          theme="auto"
          id="treeples.credit"
          routes={routes(
            this.triggerLoading.bind(this),
            this.triggerNotification.bind(this)
          )}
        >
          <Navbar>
            <div className="navbar-inner">
              <a href="/" className="link">
                Credit
              </a>
              {this.renderLoading()}
              <div className="right">{this.renderSession()}</div>
            </div>
          </Navbar>
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
          >
            {this.renderNotification()}
          </View>
        </App>
      </ErrorBoundary>
    );
  }
}
