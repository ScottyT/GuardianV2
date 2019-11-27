import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";

Vue.use(VueRouter);

const routes = [{
    path: "/",
    name: "home",
    component: () => import( /*webpackPreload: true*/ "../views/Home.vue")
  },
  {
    path: "/services",
    name: "services",
    component: () =>
      import( /*webpackPreload: true*/ "../views/Services.vue")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;