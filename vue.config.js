module.exports = {
  configureWebpack: {

  },
  chainWebpack: config => {
    ["vue-modules", "vue", "normal-modules", "normal"].forEach(match => {
      config.module
        .rule("scss")
        .oneOf(match)
        .use("sass-loader")
        .tap(opt =>
          Object.assign(opt, {
            data: `@import '~@/sass/main.scss';`
          })
        );
    });
  },
  publicPath: process.env.NODE_ENV === 'production' ?
    '/dist/' : '/',
  transpileDependencies: ["vuetify"]
};