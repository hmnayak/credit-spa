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

let navigate = null;

export const setNavigate = (fn) => {
  navigate = fn;
};

if (firebase.apps.length == 0) {
  firebase.initializeApp(config);
}

// firebase
//   .auth()
//   .getRedirectResult()
//   .then((result) => {
//     console.log("Here I am", result);
//     if (result.credential) {
//       /** @type {firebase.auth.OAuthCredential} */
//       var credential = result.credential;

//       // This gives you a Google Access Token. You can use it to access the Google API.
//       var token = credential.accessToken;
//       // ...
//     }
//     // The signed-in user info.
//     var user = result.user;
//   })
//   .catch((error) => {
//     // Handle Errors here.
//     var errorCode = error.code;
//     var errorMessage = error.message;
//     // The email of the user's account used.
//     var email = error.email;
//     // The firebase.auth.AuthCredential type that was used.
//     var credential = error.credential;
//     // ...
//   });

firebase.auth().onAuthStateChanged((curuser) => {
  console.log(curuser);
  console.log("Auth state changed");
  user = curuser;
  navigate("/");
});

export const signInWithGoogle = () => {
  const provider = new firebase.auth.GoogleAuthProvider();
  firebase.auth().signInWithRedirect(provider);
};

export const onLogoutClicked = () => {
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
    })
    .catch((error) => {
      console.error("Failed to login", error);
      alert(error.message + " Please try again.");
    });
};
