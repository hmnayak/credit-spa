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

export let user = null;

export function getFirebase() {
  let firebaseApp = !firebase.apps.length
    ? firebase.initializeApp(config)
    : firebase.app();
  return firebase;
}

export function getLoggedInUser(setUser) {
  let firebaseApp = !firebase.apps.length
    ? firebase.initializeApp(config)
    : firebase.app();

  firebase.auth().onAuthStateChanged((curuser) => {
    console.log(curuser);
    user = curuser;
    setUser(curuser);
  });
}
