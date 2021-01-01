import {
  Page,
  Block,
  Button,
  List,
  ListInput,
  BlockTitle,
} from "framework7-react";
import React from "react";
import { loginWithEmail, user } from "../auth";

export default class login extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      email: "",
      password: "",
    };
  }

  loginHeader = () => {
    return (
      <Block strong>
        <form
          onSubmit={loginWithEmail}
          action=""
          method="GET"
          className="form-ajax-submit"
        >
          <List class="login-list">
            <ListInput
              label="Login with your email"
              type="email"
              placeholder="Email address"
              value={this.state.email}
              onInput={(e) => {
                this.setState({ email: e.target.value });
              }}
            />
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
          <Button fill href="/home/" type="submit">
            Login
          </Button>
        </form>
      </Block>
    );
  };

  showSignup = () => {
    return (
      <Block>
        <BlockTitle>No Account yet?</BlockTitle>
        <Button fill href="/signup/">
          Create a new Account
        </Button>
      </Block>
    );
  };

  render() {
    if (user) {
      this.props.router("/home");
    } else {
      return (
        <Page>
          {this.loginHeader()}
          {this.showSignup()}
        </Page>
      );
    }
  }
}
