import LoginPage from "./pages/login.jsx";
import AboutPage from "./pages/about.jsx";
import HomePage from "./pages/home.jsx";
import SignupPage from "./pages/signup.jsx";
import { getCurUser } from "./services/authsvc";

export default (setLoading) => {
  return [
    {
      path: "/",
      component: HomePage,
      beforeEnter: function (router) {
        if (getCurUser() !== "Guest") {
          router.resolve();
          this.navigate("/home");
        } else {
          router.reject();
          this.navigate("/login");
        }
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
  ];
};
