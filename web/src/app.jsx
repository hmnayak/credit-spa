import React, { useState } from "react";
import { App, View, Navbar } from "framework7-react";
import LoginPage from "./pages/login.jsx";
import AboutPage from "./pages/about.jsx";
import HomePage from "./pages/home.jsx";
import SignupPage from "./pages/signup.jsx";
import { getFirebase, getLoggedInUser } from "./auth";

const rootPath = window.location.pathname.replace(/\/+$/, "");

export default () => {
  const [isLoading, setLoading] = useState(false);
  const [user, setUser] = useState(null);
  // const [firebase, setFirebase] = useState(null);

  const Loading = () => {
    if (isLoading) {
      return <div>Loading...</div>;
    }
  };

  const routes = [
    {
      path: "/",
      beforeEnter: function (router) {
        if (getLoggedInUser()) {
          setUser(getLoggedInUser());
          router.resolve();
          this.navigate("/home");
        } else {
          // setFirebase(getFirebase());
          router.reject();
          this.navigate("/login");
        }
      },
    },
    {
      path: "/login",
      component: LoginPage,
      options: {
        props: {
          setUser: setUser,
        },
      },
    },
    {
      path: "/signup",
      component: SignupPage,
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
    {
      path: "(.*)",
      component: HomePage,
      options: {
        props: {
          user: user,
        },
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
