import React, { useState, useEffect } from "react";
import { Page, Link } from "framework7-react";
import { pingApi } from "../services/pingapi";

export default (props) => {
  const [item, setItem] = useState("");

  useEffect(() => {
    pingApi(props.fetch).then((res) => {
      setItem(res.status);
    });
  }, [item]);

  return (
    <Page>
      <p>About</p>
      <div> {item}</div>
      <Link href="/">Home</Link>
    </Page>
  );
};
