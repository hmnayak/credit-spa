import { Page, Block, Button, List, ListInput } from "framework7-react";
import React from "react";
import { firebase } from "@firebase/app";
import "firebase/auth";

export default class login extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      user: null,
      email: "",
      password: "",
      emailSignup: "",
      passwordSignup: "",
      showSignupForm: true,
    };
  }

  componentDidMount() {
    const config = {
      apiKey: "AIzaSyCn3VXwkmLvubI5TZytQNH1D8nut8FoQgY",
      authDomain: "credit-7f47d.firebaseapp.com",
      projectId: "credit-7f47d",
      storageBucket: "credit-7f47d.appspot.com",
      messagingSenderId: "486648757058",
      appId: "1:486648757058:web:1232aa94de5f9be53926db",
    };

    let firebaseApp = !firebase.apps.length
      ? firebase.initializeApp(config)
      : firebase.app();

    firebase.auth().onAuthStateChanged((user) => {
      console.log(user);
      this.setState({ user: user });
    });
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
    console.log("Signup clicked");
    e.preventDefault();
    firebase
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
