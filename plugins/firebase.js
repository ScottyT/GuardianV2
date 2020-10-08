import firebase from 'firebase/app';
import 'firebase/auth';

var app;
const firebaseConfig = {
  apiKey: "AIzaSyAqDtSxeqP9oJaqcXPZGdh4xy9T3IETWL8",
  authDomain: "guardian-restoration.firebaseapp.com",
  databaseURL: "https://guardian-restoration.firebaseio.com",
  projectId: "guardian-restoration",
  storageBucket: "guardian-restoration.appspot.com",
  messagingSenderId: "684388230081",
  appId: "1:684388230081:web:b02a74ee9a3abbacb08540",
  measurementId: "G-5SJBJ1ZVDN"
};
app = firebase.apps.length ? firebase.app() : firebase.initializeApp(firebaseConfig);

export const auth = firebase.auth();
export const persistance = firebase.auth.Auth;