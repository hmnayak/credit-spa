import { Page, Block, Button, List, ListInput } from "framework7-react";
import React from "react";
import { getFirebase, user } from "../auth";
import { firebase } from "@firebase/app";

export default class login extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      user: null,
      emailSignup: "",
      passwordSignup: "",
      firebase: getFirebase(),
      showSignupForm: true,
    };
  }

  showSignup = () => {
    return (
      <Block>
        <form
          onSubmit={this.onSignupClicked}
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
          <Button fill onClick={this.onGoogleLoginClicked}>
            Sign in with Google
          </Button>
        </Block>
      </Block>
    );
  };

  render() {
    return <Page>{this.showSignup()}</Page>;
  }

  onGoogleLoginClicked() {
    const provider = new firebase.auth.GoogleAuthProvider();
    firebase.auth().signInWithRedirect(provider);
  }

  onLogoutClicked() {
    firebase
      .auth()
      .signOut()
      .catch((error) => {
        console.error("Error while trying out user", error);
      });
  }

  onSignupClicked = (e) => {
    let firebas = this.state.firebase;
    console.log("Signup clicked");
    e.preventDefault();
    firebas
      .auth()
      .createUserWithEmailAndPassword(
        this.state.emailSignup,
        this.state.passwordSignup
      )
      .catch((error) => {
        console.error("Failed to create User", error);
        alert(error.message + " Please try again", "");
      });
  };
}
