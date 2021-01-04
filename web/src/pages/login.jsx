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
    return (
      <Page>
        {this.loginHeader()}
        {this.showSignup()}
      </Page>
    );
  }

  showError = (error) => {
    console.error("Failed to login", error);
    alert(error.message + " Please try again.");
  };

  onLoginWithEmailClicked = (e) => {
    e.preventDefault();
    loginWithEmail(this.state.email, this.state.password, this.showError);
  };
}
