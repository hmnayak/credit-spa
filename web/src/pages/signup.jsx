import { Page, Block, Button, List, ListInput } from "framework7-react";
import React from "react";
import { signInWithGoogle, signUpWithEmail } from "../services/authsvc";

export default class login extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      emailSignup: "",
      passwordSignup: "",
      showSignupForm: true,
    };
  }

  showSignup = () => {
    return (
      <Page>
        <Block>
          <form
            onSubmit={signUpWithEmail}
            action=""
            method="GET"
            className="form-ajax-submit"
          >
            <List class="login-list">
              <ListInput
                label="Signup with your email"
                type="email"
                placeholder="Email address"
                value={this.state.emailSignup}
                onInput={(e) => {
                  this.setState({ emailSignup: e.target.value });
                }}
              />
              <ListInput
                label="Signup your password"
                type="password"
                placeholder="Password"
                value={this.state.passwordSignup}
                onInput={(e) => {
                  this.setState({ passwordSignup: e.target.value });
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
}
