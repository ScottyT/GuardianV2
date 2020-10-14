const admin = require('firebase-admin')
const serviceAccount = require('./gr-serviceAccounts.json')
var adminApp;
const adminConfig = {
  credential: admin.credential.cert(serviceAccount),
  databaseURL: "https://guardian-restoration.firebaseio.com"
};
adminApp = admin.apps.length ? admin.app() : admin.initializeApp(adminConfig);

const db = admin.firestore();

exports.handler = async (event, context, callback) => {
  const ref = db.collection("projects")
  var projectData = []
  await ref.get().then((qs) => {
    qs.forEach((doc) => {
      projectData.push(doc.data())
    })
  }).catch((err) => {
    console.error(err)
  })

  return {
    statusCode: 200,
    body: JSON.stringify(projectData)
  };
}