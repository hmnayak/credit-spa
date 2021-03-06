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

export let getCurUser = () => {
  return !localStorage.getItem("user") ? "Guest" : localStorage.getItem("user");
};

export let getUserToken = () => 
  !localStorage.getItem("userToken") ? "Guest" : localStorage.getItem("userToken");

if (firebase.apps.length == 0) {
  firebase.initializeApp(config);
}

firebase.auth().onAuthStateChanged((curuser) => {
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
      result.user.updateProfile({
        displayName: name,
      });
      reNavigate();
    })
    .catch((error) => {
      showError(error);
    });
};

export const loginWithEmail = async (email, password, showError, reNavigate) => {
  await firebase.auth().signInWithEmailAndPassword(email, password).catch(err => showError(err));
  const user = firebase.auth().currentUser;
  const userToken = await user.getIdToken(true);
  if (typeof Storage !== "undefined") {
    localStorage.setItem("userToken", userToken);
    localStorage.setItem("user", user.displayName);
  }
  reNavigate();
};