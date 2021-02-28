import LoginPage from "./pages/login.jsx";
import AboutPage from "./pages/about.jsx";
import HomePage from "./pages/home.jsx";
import SignupPage from "./pages/signup.jsx";
import InfoPage from "./pages/info.jsx";
import CustomerPage from "./pages/customer.jsx"
import { getCurUser } from "./services/authsvc";

export default (setLoading, setAuthScreenLoaded, setHeaderContent) => {

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

  const isAuthPageEntered = {
    props: {
      authPageLoaded: setAuthScreenLoaded,
      updateHeader : setHeaderContent
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
      options: isAuthPageEntered,
      beforeEnter: [],
    },
    {
      path: "/signup",
      component: SignupPage,
      options: isAuthPageEntered,
      beforeEnter: [],
    },
    {
      path: "/about",
      component: AboutPage,
      options: routeOpts,
      beforeEnter: [authRedirect, loadingFilter],
    },
    {
       path: "/info",
       component: InfoPage,
       beforeEnter: [],
    },
    {
      path: "/customer",
      component: CustomerPage,
      beforeEnter: [authRedirect],
   },
  ];
};
