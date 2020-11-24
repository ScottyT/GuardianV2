export default {
  /*
   ** Nuxt rendering mode
   ** See https://nuxtjs.org/api/configuration-mode
   */
  mode: 'universal',
  /*
   ** Nuxt target
   ** See https://nuxtjs.org/api/configuration-target
   */
  target: 'static',
  /*
   ** Headers of the page
   ** See https://nuxtjs.org/api/configuration-head
   */
  head: {
    titleTemplate: '%s - Guardian Restoration',
    title: process.env.npm_package_name || '',
    meta: [{
        charset: 'utf-8'
      },
      {
        name: 'viewport',
        content: 'width=device-width, initial-scale=1, viewport-fit=cover, maximum-scale=1'
      },
      {
        hid: 'description',
        name: 'description',
        content: process.env.npm_package_description || '',
      },
    ],
    link: [{
        rel: 'icon',
        type: 'image/x-icon',
        href: '/favicon.ico'
      },
      {
        rel: 'preconnect',
        href: 'https://fonts.googleapis.com',
        crossorigin: 'true'
      },
      {
        rel: 'preconnect',
        href: 'https://static.cdn.prismic.io',
        crossorigin: 'true'
      },
      {
        rel: 'stylesheet',
        href: 'https://fonts.googleapis.com/css?family=Oswald:300,600,700&display=swap'
      },
    ],
  },
  /*
   ** Global CSS
   */
  css: [
    '@/assets/css/fonts.css'
  ],
  /*
   ** Plugins to load before mounting the App
   ** https://nuxtjs.org/guide/plugins
   */
  plugins: [
    '~/plugins/font-awesome.js',
    '~/plugins/firebase.js',
    '~/plugins/vee-validate.js',
    {
      src: '~/plugins/localStorage.js',
      ssr: false
    },
    {
      src: '@/plugins/youtube-embed.js',
      mode: 'client'
    }
  ],
  /*
   ** Auto import components
   ** See https://nuxtjs.org/api/configuration-components
   */
  components: true,
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [
    // Doc: https://github.com/nuxt-community/eslint-module
    '@nuxtjs/eslint-module',
    '@nuxtjs/vuetify',
  ],
  /*
   ** Nuxt.js modules
   */
  modules: [
    '@nuxtjs/prismic',
    ['@nuxtjs/axios', {
      //baseURL: '/.netlify/functions/'
      baseURL: 'http://localhost:8081'
    }]
  ],
  router: {
    middleware: ['redirects'],
  },
  /*
   ** vuetify module configuration
   ** https://github.com/nuxt-community/vuetify-module
   */
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    treeShake: true,
    defaultAssets: {
      font: {
        family: 'Roboto'
      }
    },
    optionsPath: './vuetify.options.js'
  },
  prismic: {
    endpoint: 'https://guardianrestoration.cdn.prismic.io/api/v2',
    linkResolver: '@/plugins/link-resolver.js',
    htmlSerializer: '@/plugins/html-serializer.js',
  },
  /*
   ** Build configuration
   ** See https://nuxtjs.org/api/configuration-build/
   */
  build: {
    transpile: ["vee-validate/dist/rules", "vue-youtube-embed"],
    extractCSS: true,
    optimization: {
      minimize: true,
      splitChunks: {
        cacheGroups: {
          styles: {
            name: 'styles',
            test: /\.(css|vue)$/,
            chunks: 'all',
            enforce: true
          }
        }
      },
      namedChunks: true
    }
  },
  render: {
    resourceHints: true,
    bundleRenderer: {
      shouldPreload: (file, type) => {
        if (type === 'script' || type === 'style') {
          return true
        }
      }
    }
  },
  generate: {
    fallback: true
  }
}