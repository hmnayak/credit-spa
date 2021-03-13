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
    localStorage.setItem("user", user.displayName);
    return await user.getIdToken();
  }).catch((err) => {
    return Promise.reject(err);
  });
};

// todo: deprecate
export let getCurUser = () => {
  return !localStorage.getItem("user") ? "Guest" : localStorage.getItem("user");
};

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
  localStorage.removeItem("user");
  window.location.reload();
};

export const signUpWithEmail = (email, password, name, showError, reNavigate) => {
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
  try {
    await firebase.auth().setPersistence(firebase.auth.Auth.Persistence.LOCAL)
    await firebase.auth().signInWithEmailAndPassword(email, password)  
    reNavigate();  
  } catch(err) {
    showError(err);
  }
};