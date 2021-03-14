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

if (firebase.apps.length == 0) {
  firebase.initializeApp(config);
}

const currentUser = new Promise((resolve, reject) => {
  firebase.auth().onAuthStateChanged(user => {
    if (user) {
      resolve(user);
    } else {
      reject('nouser');
    }
  });
});

export let getToken = async () => {
  return currentUser.then(async (user) => {
    return await user.getIdToken();
  }).catch((err) => {
    return Promise.reject(err);
  });
};

export let getCurUser = async () => {
  return currentUser.then(async (user) => {
    return user.displayName;
  }).catch((err) => {
    return "Guest";
  });
};

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
  window.location.reload();
};

export const signUpWithEmail = (email, password) => {
  return firebase
    .auth()
    .createUserWithEmailAndPassword(email, password);
};

export const loginWithEmail = (email, password) => {
  return firebase.auth().setPersistence(firebase.auth.Auth.Persistence.LOCAL)
    .then(() => {
      return firebase.auth().signInWithEmailAndPassword(email, password);
    })
};