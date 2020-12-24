import React from "react";
import { Page, Link } from "framework7-react";
import Infotext from "../api";

export default () => {
  return (
    <Page>
      <p>About</p>
      <Infotext />
      <Link href="/">Home</Link>
    </Page>
  );
};
