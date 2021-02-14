import {
  Page,
  Block,
  Button,
  List,
  ListInput,
  BlockTitle,
} from "framework7-react";
import React from "react";
import { loginWithEmail } from "../services/authsvc";
import "../../css/auth.css";

export default class login extends React.Component {
  constructor(props) {
    super(props);
    props.authPageLoaded(true);
    this.state = {
      email: "",
      password: "",
    };
  }

  render() {
    return (
      <div class="page no-toolbar no-swipeback login-screen-page">
        <div class="page-content login-screen-content auth-position">
          <div class="login-screen-title">
            <a href="/" class="link">
              Credit
            </a>
          </div>
          <Block strong>
            <form
              onSubmit={this.onLoginWithEmailClicked}
              action=""
              method="GET"
              className="form-ajax-submit"
            >
              <List>
                <ListInput
                  label="Login with your email"
                  type="email"
                  placeholder="Email address"
                  value={this.state.email}
                  onInput={(e) => {
                    this.setState({ email: e.target.value });
                  }}
                  required
                ></ListInput>
                <ListInput
                  label="Provide your password"
                  type="password"
                  placeholder="Password"
                  value={this.state.password}
                  onInput={(e) => {
                    this.setState({ password: e.target.value });
                  }}
                />
              </List>
              <Button fill type="submit">
                Login
              </Button>
              <BlockTitle>No Account yet?</BlockTitle>
              <Button fill href="/signup/">
                Create a new Account
              </Button>
            </form>
          </Block>
        </div>
      </div>
    );
  }

  showError = (error) => {
    console.error("Failed to login", error);
    alert(error.message + " Please try again.");
  };

  onLoginWithEmailClicked = (e) => {
    e.preventDefault();
    loginWithEmail(this.state.email, this.state.password, this.showError).then(
      () => {
        this.props.f7router.navigate("/");
        window.location.reload();
      }
    );
  };
}
