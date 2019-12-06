<template>
  <v-app :class="this.$route.name">
    <app-navigation :key="componentKey">
      <template v-slot:site-title>Guardian Restoration</template>
      <template v-slot:nav-items>
        <v-toolbar-items class="navigation">
          <ul id="services-list" class="v-toolbar__items">
            <li v-for="(menuLink, index) in menuLinks" :key="index" class="navigation__menu-item">
              <router-link
                :key="index"
                :to="linkResolver(menuLink.link)"
              >{{ menuLink.label[0].text }}</router-link>
            </li>
          </ul>
        </v-toolbar-items>
      </template>
    </app-navigation>
    <router-view></router-view>
    <footer-nav>
      <template v-slot:site-title>Guardian Restoration</template>
      <template v-slot:nav-items>
        <ul class="footer__items">
          <li v-for="(menuLink, index) in menuLinks" :key="index" class="footer__menu-item">
            <router-link :key="index" :to="linkResolver(menuLink.link)">{{ menuLink.label[0].text }}</router-link>
            <ul id="services-list" class="footer__menu-item--services" v-if="index === 1">
              <li v-for="service in services" :key="service.id">
                <router-link :to="linkResolver(service)">{{ service.data.service_name[0].text }}</router-link>
              </li>
            </ul>
          </li>
        </ul>
      </template>
    </footer-nav>
  </v-app>
</template>
<script>
const AppNavigation = () =>
  import(/*webpackPreload: true*/ "@/components/AppNavigation");
const FooterNav = () => import(/*webpackPreload: true */ "@/components/Footer");
export default {
  name: "App",
  components: {
    AppNavigation,
    FooterNav
  },
  data: () => ({
    menuContent: [],
    menuLinks: [],
    linkResolver: null,
    services: [],
    componentKey: 0,
    counter: 0,
    hasSideNav: false
  }),
  mounted() {},
  methods: {
    getMenu() {
      this.$prismic.client.getSingle("menu").then(menuContent => {
        this.menuContent = menuContent;
        this.menuLinks = menuContent.data.menu_links;
      });

      this.linkResolver = this.$prismic.linkResolver;
    },
    getServices() {
      this.$prismic.client
        .query(this.$prismic.Predicates.at("document.type", "services"))
        .then(services => {
          this.services = services.results;
        });
    }
  },

  computed: {},
  created() {
    this.getMenu();
    this.getServices();
  },
  beforeUpdate() {
    this.componentKey += 1;
  }
};
</script>
<style lang="scss">
.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  border: 0;
}
.v-application {
  a {
    text-decoration: none;
    color: inherit !important;
  }
}

html {
  scroll-behavior: smooth;

  @include respond(mobileSmall) {
    font-size: 50.5% !important;
  }

  @include respond(iPhoneX) {
    font-size: 50.5% !important;
  }

  @include respond(mobileLarge) {
    font-size: 54% !important;
  }

  @include respond(tabletSmall) {
    font-size: 58% !important;
  }

  @include respond(tabletLarge) {
    font-size: 62.5% !important;
  }
  @include respond(iPhoneX) {
    font-size: 54% !important;
  }
  @include respond(iPhoneXR) {
    font-size: 54% !important;
  }
}

body {
  font-size: 1.8rem;
}

html,
body {
  height: 100%;
}

h1 {
  font-family: $heading-font;
  letter-spacing: 0rem;
  font-size: 8rem;
  color: $primary-dark;

  @include respond(mobileSmall) {
    letter-spacing: 1px;
    font-size: 6rem;
    line-height: 7rem;
  }

  &.extra-large {
    font-weight: 600;
    margin-bottom: 2rem;

    @include respond(tabletSmall) {
      font-size: 9rem;
      line-height: 10rem;
    }

    @include respond(tabletLarge) {
      font-size: 9rem;
      line-height: 10rem;
    }

    @include respond(iPhoneX) {
      font-size: 5rem;
      line-height: 5rem;
    }
    @include respond(iPhoneXR) {
      font-size: 5rem;
      line-height: 5rem;
    }
    @include respond(GalaxyS10Plus) {
      font-size: 5rem;
      line-height: 5rem;
    }
    @include respond(desktopSmall) {
      font-size: 11rem;
      line-height: 13rem;
    }
  }
}
h2 {
  font-family: $heading-font;
  letter-spacing: 0.1rem;

  @include respond(mobileSmall) {
    font-size: 5rem;
    line-height: 5.5rem;
  }

  @media only screen and (max-width: 900px) {
    letter-spacing: 0;
  }

  @include respond(mobileLand) {
    font-size: 4rem;
    line-height: 5rem;
  }

  @include respond(mobilePort) {
    font-size: 5rem;
    line-height: 5.5rem;
  }

  @include respond(tabletSmall) {
    font-size: 5rem;
    line-height: 5.5rem;
  }
  @include respond(tabletLand) {
    font-size: 4.5rem;
    line-height: 5rem;
  }
  @include respond(desktopSmall) {
    font-size: 6rem;
    line-height: 6.5rem;
  }

  @include respond(iPhoneX) {
    font-size: 4rem;
    line-height: 4.5rem;
  }
  @include respond(iPhoneXR) {
    font-size: 4rem;
    line-height: 4.5rem;
  }
}

p {
  &.subheading {
    font-size: 2.5rem;
    font-weight: 600;
    text-transform: uppercase;
    margin-bottom: 3rem;

    @include respond(mobileSmall) {
      font-size: 2rem;
      margin-bottom: 2rem;
    }

    @include respond(mobileSmallLand) {
      font-size: 2rem;
      margin-bottom: 1.5rem;
    }

    @include respond(iPhoneX) {
      margin-bottom: 1rem;
    }
  }

  &.light {
    font-weight: 300;
  }
}
ul {
  list-style: none;
}
.no-padding {
  padding: 0 !important;
}
hr {
  height: 2px;
  border: 0;
  &.small {
    background-color: $secondary-dark;
    width: 50px;
    margin: 0 auto 3rem auto;
  }
}
</style>
