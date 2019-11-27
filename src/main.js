import Vue from "vue";
const App = () => import( /*webpackPreload: true*/ "./App.vue");
import PrismicVue from "prismic-vue";
import linkResolver from "./prismic/link-resolver";
import htmlSerializer from "./prismic/html-serializer";
import router from "./router";
import vuetify from "./plugins/vuetify";

Vue.config.productionTip = false;

Vue.use(PrismicVue, {
  endpoint: window.prismic.endpoint,
  linkResolver,
  htmlSerializer
});

Vue.filter("allCaps", function (value) {
  if (!value) return "";
  value = value.toString();
  return value.toUpperCase();
});

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount("#app");