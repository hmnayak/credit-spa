import {
  Page,
  Block,
  Button,
  List,
  ListInput,
  BlockTitle,
} from "framework7-react";
import React from "react";
// import { firebase } from "@firebase/app";
// import "firebase/auth";
import { getFirebase, user } from "../auth";

export default class login extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      email: "",
      password: "",
      firebase: getFirebase(),
      updateUser: props.setUser,
    };
    console.log(this.state.firebase);
  }

  loginHeader = () => {
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
      this.props.history("/home");
    } else {
      return (
        <Page>
          {this.loginHeader()}
          {this.showSignup()}
        </Page>
      );
    }
  }

  onLoginWithEmailClicked = (e) => {
    e.preventDefault();
    // const appIns = this.props.f7router.app;
    // console.log(appIns);
    let firebase = this.state.firebase;
    // console.log(firebase);
    firebase
      .auth()
      .signInWithEmailAndPassword(this.state.email, this.state.password)
      .then((res) => {
        // console.log(firebase.auth().currentUser);
        // console.log(this.state.updateUser);
        this.state.updateUser(firebase.auth().currentUser);
        // this.state.navigate("/home/");
      })
      .catch((error) => {
        console.error("Failed to login", error);
        alert(error.message + " Please try again.");
      });
    this.setState({ firebase: firebase });
  };

  onLogoutClicked() {
    this.state.firebase
      .auth()
      .signOut()
      .catch((error) => {
        console.error("Error while trying out user", error);
      });
  }
}
