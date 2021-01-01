import LoginPage from "./pages/login.jsx";
import AboutPage from "./pages/about.jsx";
import HomePage from "./pages/home.jsx";
import SignupPage from "./pages/signup.jsx";

export default (setLoading) => {
  return [
    {
      path: "/",
      component: HomePage,
      beforeEnter: function (router) {
        router.resolve();
        // if (user) {
        //   router.resolve();
        //   this.navigate("/home");
        // } else {
        //   router.reject();
        //   this.navigate("/login");
        // }
      },
    },
    {
      path: "/login",
      component: LoginPage,
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
    // {
    //   path: "(.*)",
    //   component: HomePage,
    //   options: {
    //     props: {
    //       user: user,
    //     },
    //   },
    // },
  ];
};
