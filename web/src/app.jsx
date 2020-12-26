import React, { useState } from "react";
import { App, View, Navbar } from "framework7-react";
import HomePage from "./pages/home.jsx";
import AboutPage from "./pages/about.jsx";

const rootPath = window.location.pathname.replace(/\/+$/, "");

export default () => {
  const [isLoading, setLoading] = useState(false);

  const Loading = () => {
    if (isLoading) {
      return <div>Loading...</div>;
    }
  };

  const routes = [
    {
      path: "/",
      component: HomePage,
    },
    {
      path: "/about",
      component: AboutPage,
      options: {
        props: {
          loadComplete: setLoading,
        },
      },
      beforeEnter: function (router) {
        setLoading(true);
        router.resolve();
      },
    },
  ];

  return (
    <App name="Credit" theme="auto" id="treeples.credit" routes={routes}>
      <Navbar title="Credit">{Loading()}</Navbar>
      <View
        main
        url={rootPath}
        browserHistory
        browserHistorySeparator=""
        browserHistoryRoot=""
        animate={false}
      />
    </App>
  );
};
