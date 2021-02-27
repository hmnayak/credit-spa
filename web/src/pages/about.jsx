import React, { useState, useEffect } from "react";
import { Page, Link } from "framework7-react";
import { aboutInfoApi } from "../services/api";

export default (props) => {
  const [item, setItem] = useState("");

  let aboutPromise = aboutInfoApi();

  useEffect(() => {
    console.log("Inside useEffect");
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
      <Link href="/">Home</Link>
    </Page>
  );
};
