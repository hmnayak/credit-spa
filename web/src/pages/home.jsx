import React from "react";
import { Page, Link, Block, Button } from "framework7-react";
import { logoutClicked, getCurUser } from "../services/authsvc";

export default (props) => {
  const userInfo = () => {
    return <Block strong>You are logged in as {getCurUser()}</Block>;
  };

  return (
    <Page>
      <Block strong>
        <p>Life's made simple!</p>
      </Block>
      {userInfo()}
    </Page>
  );
};
