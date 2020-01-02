const BundleAnalyzerPlugin = require('webpack-bundle-analyzer').BundleAnalyzerPlugin;
const PreloadPlugin = require("@vue/preload-webpack-plugin");

module.exports = {
  configureWebpack: {
    plugins: [
      // new BundleAnalyzerPlugin(),
      new PreloadPlugin({
        rel: "preload",
        include: "allAssets",
        as(entry) {
          if (/\.css$/.test(entry)) return "style";
          return "script";
        }
      })
    ],
    performance: {
      maxEntrypointSize: 812000,
      maxAssetSize: 812000
    },
    output: {
      chunkFilename: "[name].bundle.js"
    },
    // resolve: {
    //   alias: {
    //     vue$: "vue/dist/vue.esm.js" // 'vue/dist/vue.common.js' for webpack 1
    //   }
    // },
    optimization: {
      splitChunks: {
        chunks: "all"
      }
    }
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
    config.plugins.delete("prefetch");
  },

  transpileDependencies: ["vuetify"]
};