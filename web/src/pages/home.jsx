import React from "react";
import { Page, Link, Block, Button } from "framework7-react";
import { logoutClicked } from "../services/authsvc";

export default (props) => {
  const userInfo = () => {
    let user = null;
    if (localStorage.getItem("user")) {
      user = localStorage.getItem("user");
    } else {
      user = "Guest";
    }
    return (
      <Block strong>
        You are logged in as {user}
        <Button onClick={logoutClicked}>Logout</Button>
      </Block>
    );
  };

  return (
    <Page>
      <Block strong>
        <p>Hello world</p>
        <Link href="/about/">About</Link>
        <br />
        <Link href="/login/">Login</Link>
        <br />
        <Link href="/signup/">Signup</Link>
        <br />
      </Block>
      {userInfo()}
    </Page>
  );
};
