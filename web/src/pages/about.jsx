import React, { useState, useEffect } from "react";
import { Page, Link } from "framework7-react";
import { aboutinfoApi } from "../api";

export default (props) => {
  console.log(props);
  const [item, setItem] = useState("");

  let aboutPromise = aboutinfoApi();

  useEffect(() => {
    aboutPromise
      .then((res) => res.text())
      .then((result) => {
        props.loadComplete(false);
        setItem(result);
      });
  }, []);

  return (
    <Page>
      <p>About</p>
      <div> {item}</div>
      <Link href=".*">Home</Link>
    </Page>
  );
};
