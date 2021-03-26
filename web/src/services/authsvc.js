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

const userChangeCallbacks = [];

export const onUserChange = (fn) => {
  userChangeCallbacks.push(fn);
}

const currentUser = new Promise((resolve, reject) => {
    firebase.auth().onAuthStateChanged(user => {
      if (user) {
        resolve(user);
      } else {
        reject('nouser');
      }
      userChangeCallbacks.forEach(fn => fn());
    });
  });

export const getToken = async () => {
  return currentUser.then(async (user) => {
    return await user.getIdToken();
  }).catch((err) => {
    return Promise.reject(err);
  });
};

export const getUsername = () => {
  return currentUser.then(user => {
    return user.displayName;
  }).catch(() => {
    return null;
  });
};

export const signUpWithEmail = (email, password) => {
  return firebase
    .auth()
    .createUserWithEmailAndPassword(email, password);
};

export const signInWithGoogle = () => {
  const provider = new firebase.auth.GoogleAuthProvider();
  return firebase.auth().signInWithRedirect(provider);
};

export const loginWithEmail = (email, password) => {
  return firebase.auth()
    .setPersistence(firebase.auth.Auth.Persistence.LOCAL)
    .then(() => {
      return firebase.auth().signInWithEmailAndPassword(email, password);
    });
}

export const logoutClicked = () => {
  return firebase
    .auth()
    .signOut()
    .catch((error) => {
      console.error("Error while trying out user", error);
    });
};
