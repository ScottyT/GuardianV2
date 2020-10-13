export default function(ctx) {
  if (ctx.route.fullPath == '/user' || ctx.route.fullPath == '/user/') {
    ctx.redirect(301, '/user/projects')
  }
}