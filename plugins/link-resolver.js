/**
 * To learn more about Link Resolving check out the Prismic documentation
 * https://prismic.io/docs/vuejs/beyond-the-api/link-resolving
 */

export default function (doc) {
  if (doc.isBroken) {
    return '/not-found'
  }

  if (doc.type === 'home') {
    return '/'
  }

  if (doc.type === "service") {
    return "/service";
  }
  if (doc.type === "services") {
    return "/service/" + doc.uid;
  }
  if (doc.type === "page") {
    return "/" + doc.uid;
  }
  if (doc.type === 'user') {
    return "/user/" + doc.uid;
  }
  return '/' + doc.type + '/' + doc.uid
}
