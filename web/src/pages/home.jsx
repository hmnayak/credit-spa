import React from "react";
import { Page, Link, Block } from "framework7-react";
import { user } from "../services/authsvc";

export default (props) => {
  const userInfo = () => {
    return (
      <Block strong>
        <Block-header>
          You are logged in as {user.displayName} {user.email}
        </Block-header>
        <Button onClick={this.onLogoutClicked}>Logout</Button>
      </Block>
    );
  };

  return (
    <Page>
      {/* {userInfo()} */}
      <p>Hello world</p>
      <Link href="/about/">About</Link>
      <br />
      <Link href="/login/">Login</Link>
      <br />
      <Link href="/signup/">Signup</Link>
      <br />
    </Page>
  );
};
