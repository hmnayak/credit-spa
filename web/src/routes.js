import LoginPage from "./pages/login.jsx";
import AboutPage from "./pages/about.jsx";
import HomePage from "./pages/home.jsx";
import SignupPage from "./pages/signup.jsx";
import NewCustomerPage from "./pages/customers/customer.jsx"
import CustomersPage from "./pages/customers.jsx"
import { fetchFn as fetchFn } from "./services/api";

export default (setLoading, setAuthScreenLoaded, setHeaderContent) => {

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
      options: {
        props : {
          fetch: fetchFn(setLoading),
          loadComplete: setLoading,
        }
      },
      beforeEnter: [],
    },
    {
      path: "/customers",
      component: CustomersPage,
      options: {
        props : {
          fetch: fetchFn(setLoading),
          loadComplete: setLoading,
        }
      },
      beforeEnter: [],
    },
    {
      path: "/customers/new",
      component: NewCustomerPage,
      options: {
        props : {
          fetch: fetchFn(setLoading, true),
          loadComplete: setLoading,
        }
      },
      beforeEnter: [],
    },
  ];
};
