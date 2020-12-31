import React from "react";
import { Page, Link } from "framework7-react";

export default (props) => {
  const userInfo = () => {
    <Block strong>
      <Block-header>
        You are logged in as {props.user.displayName} {props.user.email}
      </Block-header>
      <Button onClick={this.onLogoutClicked}>Logout</Button>
    </Block>;
  };
  return (
    <Page>
      <p>Hello world</p>
      {userInfo}
      <Link href="/about/">About</Link>
    </Page>
  );
};
