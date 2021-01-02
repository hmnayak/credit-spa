import { Page, Block, Button, List, ListInput } from "framework7-react";
import React from "react";
import { signInWithGoogle, signUpWithEmail } from "../services/authsvc";

export default class login extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      email: "",
      password: "",
    };
  }

  showSignup = () => {
    return (
      <Page>
        <Block>
          <form
            onSubmit={this.onSignupWithEmailClicked.bind(this)}
            action=""
            method="GET"
            className="form-ajax-submit"
          >
            <List class="login-list">
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
            </List>
            <Button fill type="submit">
              Signup
            </Button>
          </form>
          <Block>
            <Button fill onClick={signInWithGoogle}>
              Sign in with Google
            </Button>
          </Block>
        </Block>
      </Page>
    );
  };

  render() {
    return <Page>{this.showSignup()}</Page>;
  }

  onSignupWithEmailClicked = (e) => {
    e.preventDefault();
    signUpWithEmail(this.state.email, this.state.password);
  };
}