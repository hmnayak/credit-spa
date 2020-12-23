import React, { useState, useEffect } from "react";
import { Page, Link } from "framework7-react";

export default () => (
  <Page>
    <p>About</p>
    <Aboutinfo />
    <Link href="/">Home</Link>
  </Page>
);

const Aboutinfo = (props) => {
  const [error, setError] = useState(null);
  const [isLoaded, setIsLoaded] = useState(false);
  const [items, setItem] = useState("");

  // Note: the empty deps array [] means
  // this useEffect will run once
  // similar to componentDidMount()
  useEffect(() => {
    fetch("/api/ping")
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
    return <div>Loading...</div>;
  } else {
    return <div> {items}</div>;
  }
};
