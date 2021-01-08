import LoginPage from "./pages/login.jsx";
import AboutPage from "./pages/about.jsx";
import HomePage from "./pages/home.jsx";
import SignupPage from "./pages/signup.jsx";
import { getCurUser } from "./services/authsvc";

export default (setLoading) => {

  const authRedirect = (routeContext) => {
    if (getCurUser() !== "Guest") {
      routeContext.resolve();
    } else {
      routeContext.reject();
      routeContext.router.navigate('/login');
    }
  }

  const loadingFilter = (routeContext) => {
    setLoading(true);
    routeContext.resolve();
  }

  const routeOpts = {
    props: {
      loadComplete: setLoading,
    },
  }

  return [
    {
      path: "/",
      component: HomePage,
      beforeEnter: [],
    },
    {
      path: "/home",
      component: HomePage,
    },
    {
      path: "/login",
      component: LoginPage,
      beforeEnter: [],
    },
    {
      path: "/signup",
      component: SignupPage,
      beforeEnter: [],
    },
    {
      path: "/about",
      component: AboutPage,
      options: routeOpts,
      beforeEnter: [authRedirect, loadingFilter],
    },
  ];
};
