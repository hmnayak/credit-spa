import React from "react";
import { Page, Link } from "framework7-react";
import AboutInfo from "./aboutinfo";

export default () => (
  <Page>
    <p>About</p>
    <AboutInfo />
    <Link href="/">Home</Link>
  </Page>
);
