import { firebase } from "@firebase/app";
import "firebase/auth";

const config = {
  apiKey: "AIzaSyCn3VXwkmLvubI5TZytQNH1D8nut8FoQgY",
  authDomain: "credit-7f47d.firebaseapp.com",
  projectId: "credit-7f47d",
  storageBucket: "credit-7f47d.appspot.com",
  messagingSenderId: "486648757058",
  appId: "1:486648757058:web:1232aa94de5f9be53926db",
};

let navigate = null;

let user = null;

export let getCurUser = () => {
  return !localStorage.getItem("user") ? "Guest" : localStorage.getItem("user");
};

export const setNavigate = (fn) => {
  navigate = fn;
};

if (firebase.apps.length == 0) {
  firebase.initializeApp(config);
}

firebase.auth().onAuthStateChanged((curuser) => {
  user = curuser;
  if (navigate) {
    navigate("/");
  }
});

export const signInWithGoogle = () => {
  const provider = new firebase.auth.GoogleAuthProvider();
  firebase.auth().signInWithRedirect(provider);
};

export const logoutClicked = () => {
  localStorage.setItem("user", "Guest");
  firebase
    .auth()
    .signOut()
    .catch((error) => {
      console.error("Error while trying out user", error);
    });
  window.location.reload();
};

export const signUpWithEmail = (email, password, name, showError) => {
  firebase
    .auth()
    .createUserWithEmailAndPassword(email, password)
    .then((result) => {
      return result.user.updateProfile({
        displayName: name,
      });
    })
    .catch((error) => {
      showError(error);
    });
};

export const loginWithEmail = (email, password, showError) => {
  firebase
    .auth()
    .signInWithEmailAndPassword(email, password)
    .then((res) => {
      user = firebase.auth().currentUser;
      if (typeof Storage !== "undefined") {
        localStorage.setItem("user", user.displayName);
      }
    })
    .catch((error) => {
      showError(error);
    });
};
