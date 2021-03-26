import { Block, Button, List, ListInput } from "framework7-react";
import React from "react";
import { signInWithGoogle, signUpWithEmail } from "../services/authsvc";
import "../../css/auth.css";

export class SignupPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      email: "",
      password: "",
      name: "",
      errorMsg: "",
    };
  }

  render() {
    return (
      <div className="page no-toolbar no-swipeback login-screen-page">
        <div className="page-content login-screen-content auth-position">
          <div className="login-screen-title">
            <a href="/" className="link">
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
  }

  reNavigate() {
    window.location.href = '/';
  }

  onSignupWithEmailClicked(e) {
    e.preventDefault();
    signUpWithEmail(this.state.email, this.state.password)
      .then((result) => {
        result.user.updateProfile({
          displayName: this.state.name,
        });
        this.reNavigate();
      })
      .catch((error) => {
        this.showError(error);
      });
  }
}
