import React, { useState, useEffect } from "react";
import { Page, Link } from "framework7-react";
import { aboutInfoApi } from "../services/api";

export default (props) => {
  const [item, setItem] = useState("");
  const [aboutPromise, setAboutPromise] = useState(null);

  function receivePromise(prom) {
    setAboutPromise(prom);
  }

  if (!aboutPromise) {
    aboutInfoApi(receivePromise);
  }

  if (aboutPromise) {
    aboutPromise
      .then((res) => res.text())
      .then((result) => {
        props.loadComplete(false);
        setItem(result);
      });
  }

  return (
    <Page>
      <p>About</p>
      <div> {item}</div>
      <Link href="/">Home</Link>
    </Page>
  );
};
