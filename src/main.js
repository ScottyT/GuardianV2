import Vue from "vue";
const App = () => import("./App.vue");
import PrismicVue from "prismic-vue";
import linkResolver from "./prismic/link-resolver";
import htmlSerializer from "./prismic/html-serializer";
import router from "./router";
import vuetify from "./plugins/vuetify";

Vue.config.productionTip = false;

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount("#app");