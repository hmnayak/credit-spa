import React, { useState, useEffect } from "react";
import { Page, Link } from "framework7-react";
import { aboutinfoApi } from "../api";
import { Loading } from "../app.jsx";

export default () => {
  return (
    <Page>
      <p>About</p>
      <Infotext />
      <Link href="/">Home</Link>
    </Page>
  );
};

const Infotext = (props) => {
  const [error, setError] = useState(null);
  const [isLoaded, setIsLoaded] = useState(false);
  const [item, setItem] = useState("");

  let aboutPromise = aboutinfoApi();

  // Note: the empty deps array [] means
  // this useEffect will run once
  // similar to componentDidMount()
  useEffect(() => {
    aboutPromise
      .then((res) => res.text())
      .then(
        (result) => {
          setIsLoaded(true);
          console.log(result);
          setItem(result);
        },
        // Note: it's important to handle errors here
        // instead of a catch() block so that we don't swallow
        // exceptions from actual bugs in components.
        (error) => {
          setIsLoaded(true);
          setError(error);
        }
      );
  }, []);

  if (error) {
    return <div>Error: {error.message}</div>;
  } else if (!isLoaded) {
    return Loading();
  } else {
    return <div> {item}</div>;
  }
};
