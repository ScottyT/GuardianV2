const admin = require('firebase-admin')
const fs = require('fs');
const path = require('path');

const serviceAccount = require('./gr-serviceAccounts.json');
var adminApp;
const adminConfig = {
  credential: admin.credential.cert(serviceAccount),
  databaseURL: "https://guardian-restoration.firebaseio.com"
};
adminApp = admin.apps.length ? admin.app() : admin.initializeApp(adminConfig);

const db = admin.firestore();
const auth = admin.auth();

exports.handler = async (event, context) => {
  const { id } = JSON.parse(event.body)
  console.log(id)
  const ref = db.collection("signed-in-users");
  var signedIn = ""

  await ref.doc(id).get().then((doc) => {
    if (doc.exists) {
      signedIn = "true"
    } else {
      console.log("false")
      ref.doc(id).set({
        id: id
      })
      signedIn = "false"
    }
  })
  // const { email, id } = JSON.parse(event.body)
  // var data = '';
  // var success = '';
  // const userRef = db.collection('users').doc(email)
  // await userRef.get().then((doc) => {
  //   if (doc.exists) {
  //     data = doc.data()
  //   } else {
  //     userRef.set({
  //       email: email,
  //       id: id,
  //       role: "user"
  //     }).then(() => {
  //       success = "created user"
  //     }).catch(() => {
  //       console.log("no user create")
  //     })
  //   }
  // })
  return {
    statusCode: 200,
    body: signedIn
  }
}