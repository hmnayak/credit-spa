import { LoginPage } from "./pages/login.jsx";
import { AboutPage } from "./pages/about.jsx";
import { HomePage } from "./pages/home.jsx";
import { SignupPage } from "./pages/signup.jsx";
import { CustomerPage as CustomerPage } from "./pages/customers/customer.jsx"
import { ListCustomersPage } from "./pages/customers.jsx"
import { fetchFn } from "./services/api";

export default (setLoading, showNotification) => {

  const routeOpts  = {
    props : {
      fetch: fetchFn(setLoading),
      loadComplete: setLoading,
      showNotification: showNotification
    }
  }

  return [
    {
      path: "",
      component: HomePage,
      options: {
        props : {
        }
      },
      beforeEnter: [],
    },
    {
      path: "/home",
      component: HomePage,
      options: {
        props : {
        }
      },
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
      options: {
        props : {
          fetch: fetchFn(setLoading),
          loadComplete: setLoading,
        }
      },
      beforeEnter: [],
    },
    {
      path: "/customers/",
      component: ListCustomersPage,
      options: routeOpts,
      beforeEnter: [],
    },
    {
      path: "/customers/\\?page=(.*)",
      component: ListCustomersPage,
      options: routeOpts,
      beforeEnter: [],
    },
    {
      path: "/customers/new",
      component: CustomerPage,
      options: routeOpts,
      beforeEnter: [],
    },
    {
      path: "/customers/:customerId",
      component: CustomerPage,
      options: routeOpts,
      beforeEnter: [],
    },
  ];
};
