import { Page, Block, Button, List, ListInput } from "framework7-react";
import React from "react";
import { signInWithGoogle, signUpWithEmail } from "../services/authsvc";
import "../../css/auth.css";

export default class signup extends React.Component {
  constructor(props) {
    super(props);
    props.authPageLoaded(true);
    this.state = {
      email: "",
      password: "",
      name: "",
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
              onSubmit={this.onSignupWithEmailClicked}
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
        </div>
      </div>
    );
  }

  showError = (error) => {
    console.error("Failed to create User", error);
    alert(error.message + " Please try again", "");
  };

  reNavigate = () => {
    this.props.f7router.navigate("/");
    window.location.reload();
  };

  onSignupWithEmailClicked = (e) => {
    e.preventDefault();
    signUpWithEmail(
      this.state.email,
      this.state.password,
      this.state.name,
      this.showError,
      this.reNavigate
    );
  };
}
