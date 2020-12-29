import {
  Page,
  Block,
  Button,
  List,
  ListInput,
  BlockTitle,
} from "framework7-react";
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
    // firebase.initializeApp(config);

    firebase.auth().onAuthStateChanged((user) => {
      console.log(user);
      this.state.user = user;
      this.setState({ user: user });
    });
  }

  loginHeader = () => {
    console.log(this.state.user);
    if (this.state.user) {
      return (
        <Block strong>
          <Block-header>
            You are logged in as {this.state.user.displayName}{" "}
            {this.state.user.email}
          </Block-header>
          <Button onClick={this.onLogoutClicked}>Logout</Button>
        </Block>
      );
    } else {
      return (
        <Block strong>
          <form
            onSubmit={this.onLoginWithEmailClicked.bind(this)}
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
            <Button fill type="submit">
              Login
            </Button>
          </form>
        </Block>
      );
    }
  };

  updateShowSignUp() {
    console.log("Show Signup");
    this.setState({
      showSignupForm: false,
    });
  }

  showSignup = () => {
    if (this.state.showSignupForm) {
      return (
        <Block>
          <BlockTitle>No Account yet?</BlockTitle>
          <Button fill onClick={this.updateShowSignUp.bind(this)}>
            Create a new Account
          </Button>
        </Block>
      );
    } else {
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
    }
  };

  render() {
    return (
      <Page>
        {this.loginHeader()}
        {this.showSignup()}
      </Page>
    );
  }

  onLoginWithEmailClicked = (e) => {
    e.preventDefault();
    const appIns = this.props.f7router.app.params;
    firebase
      .auth()
      .signInWithEmailAndPassword(this.state.email, this.state.password)
      .then((res) => {
        this.state.user = firebase.auth().currentUser.user;
        this.setState({ user: firebase.auth().currentUser.user });
        console.log(this.state.user);
      })
      .catch((error) => {
        console.error("Failed to login", error);
        alert(error.message + " Please try again.");
      });
  };

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
