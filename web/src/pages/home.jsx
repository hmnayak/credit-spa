import React from "react";
import { Page, Link, Block, Button } from "framework7-react";
import { user, logoutClicked } from "../services/authsvc";

export default (props) => {
  const userInfo = () => {
    console.log(user);
    if (user) {
      return (
        <Block strong>
          <Block-header>
            You are logged in as {user.displayName} {user.email}
            <Button onClick={logoutClicked}>Logout</Button>
          </Block-header>
        </Block>
      );
    }
  };

  return (
    <Page>
      {userInfo()}
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
