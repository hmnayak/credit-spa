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

export let user = null;

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
  firebase
    .auth()
    .signOut()
    .catch((error) => {
      console.error("Error while trying out user", error);
    });
};

export const signUpWithEmail = (email, password) => {
  firebase
    .auth()
    .createUserWithEmailAndPassword(email, password)
    .catch((error) => {
      console.error("Failed to create User", error);
      alert(error.message + " Please try again", "");
    });
};

export const loginWithEmail = (email, password) => {
  firebase
    .auth()
    .signInWithEmailAndPassword(email, password)
    .then((res) => {
      user = firebase.auth().currentUser;
      console.log(user);
      navigate("/");
    })
    .catch((error) => {
      console.error("Failed to login", error);
      alert(error.message + " Please try again.");
    });
};
