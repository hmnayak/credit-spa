import LoginPage from "./pages/login.jsx";
import AboutPage from "./pages/about.jsx";
import HomePage from "./pages/home.jsx";
import SignupPage from "./pages/signup.jsx";
import NewCustomerPage from "./pages/customers/customer.jsx"
import CustomersPage from "./pages/customers.jsx"
import { getCurUser } from "./services/authsvc";
import { fetchFnJson, fetchFn } from "./services/api";

export default (setLoading, setAuthScreenLoaded, setHeaderContent) => {

  // todo: deprecate
  const authRedirect = (routeContext) => {
    if (getCurUser() !== "Guest") {
      routeContext.resolve();
    } else {
      routeContext.reject();
      routeContext.router.navigate('/login');
    }
  }

  // todo: deprecate
  const loadingFilter = (routeContext) => {
    setLoading(true);
    routeContext.resolve();
  }

  const routeOpts = {
    props: {
      fetch: fetchFnJson(setLoading),
      loadComplete: setLoading,
    },
  }

  const routeOptsCust = {
    props: {
      fetch: fetchFn(),
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
      beforeEnter: [],
    },
    {
      path: "/customers",
      component: CustomersPage,
      options: routeOpts,
      beforeEnter: [],
    },
    {
      path: "/customers/new",
      component: NewCustomerPage,
      options: routeOptsCust,
      beforeEnter: [],
    },
  ];
};
