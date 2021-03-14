import { Block, Button, List, ListInput } from "framework7-react";
import React from "react";
import { signInWithGoogle, signUpWithEmail } from "../services/authsvc";
import "../../css/auth.css";

export class SignupPage extends React.Component {
  constructor(props) {
    super(props);
    props.authPageLoaded(true);
    this.state = {
      email: "",
      password: "",
      name: "",
      errorMsg: "",
    };
  }

  onHomeLinkClicked() {
    this.props.authPageLoaded(false);
    this.props.updateHeader();
  };

  render() {
    return (
      <div className="page no-toolbar no-swipeback login-screen-page">
        <div className="page-content login-screen-content auth-position">
          <div className="login-screen-title">
            <a href="/" onClick={this.onHomeLinkClicked.bind(this)} className="link">
              Credit
            </a>
          </div>
          <Block>
            <form
              onSubmit={this.onSignupWithEmailClicked.bind(this)}
              action=""
              method="GET"
              className="form-ajax-submit"
            >
              <List className="login-list">
                <ListInput
                  label="Name"
                  type="text"
                  placeholder="Name"
                  value={this.state.name}
                  onInput={(e) => {
                    this.setState({ name: e.target.value });
                  }}
                />
                <ListInput
                  label="Signup with your email"
                  type="email"
                  placeholder="Email address"
                  value={this.state.email}
                  onInput={(e) => {
                    this.setState({ email: e.target.value });
                  }}
                />
                <ListInput
                  label="Signup your password"
                  type="password"
                  placeholder="Password"
                  value={this.state.password}
                  onInput={(e) => {
                    this.setState({ password: e.target.value });
                  }}
                />
                <p style={{ color: "red" }}>{this.state.errorMsg}</p>
              </List>
              <Button fill type="submit">
                Signup
              </Button>
            </form>
            <Block>
              <Button fill onClick={signInWithGoogle.bind(this)}>
                Sign in with Google
              </Button>
            </Block>
          </Block>
        </div>
      </div>
    );
  }

  showError(error) {
    this.setState({
      errorMsg: error.message,
    });
  };

  reNavigate() {
    this.onHomeLinkClicked();
    this.props.f7router.navigate("/");
  };

  onSignupWithEmailClicked(e) {
    e.preventDefault();
    signUpWithEmail(
      this.state.email,
      this.state.password,
      this.state.name,
      this.showError.bind(this),
      this.reNavigate.bind(this)
    );
  };
}
