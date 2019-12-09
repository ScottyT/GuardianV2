<template>
  <div class="content services-page" v-if="loaded">
    <div class="row">
      <div id="log"></div>
      <div class="services-page__intro" @click="titleSelect($event)">
        <h1 refs="intro" class="services-page__intro--title">Services</h1>
        <hr class="small" />
      </div>
      <div class="services-page__columns">
        <v-col
          class="service-pages service-item"
          v-for="service in oddServices"
          :key="service.id"
          cols="12"
          lg="6"
          @click="serviceSelect($event)"
          @mouseenter="mouseEnter($event)"
          @mouseleave="mouseLeave($event)"
        >
          <v-card class="service-item__card" elevation="10">
            <font-awesome-layers class="fa-3x service-item__icon">
              <font-awesome-icon class="service-item__icon--background" :icon="['fas', 'circle']" />
              <font-awesome-icon
                class="service-item__icon--plus"
                :icon="['fas', 'plus']"
                transform="shrink-6"
              />
            </font-awesome-layers>

            <v-img
              src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
              class="service-item__background white--text"
              gradient="120deg, rgba(0,0,0,0.88) 0%, rgba(0,0,0,0.6) 45%, rgba(0,0,0,0.4) 60%,  rgba(0,0,0,0) 100%"
            >
              <h2 class="service-item__title">{{ service.title }}</h2>
              <p class="light">{{ service.serviceText }}</p>
            </v-img>
            <div class="service-item__big-description">
              <div class="service-item__big-description--wrap">
                <p>{{ service.bigDescription }}</p>
              </div>
            </div>
          </v-card>
        </v-col>
      </div>
      <div class="services-page__columns">
        <v-col
          class="service-pages service-item"
          v-for="service in evenServices"
          :key="service.id"
          cols="12"
          lg="6"
          @click="serviceSelect($event)"
          :refs="'service--' + service.title"
          @mouseenter="mouseEnter($event)"
          @mouseleave="mouseLeave($event)"
        >
          <v-card class="service-item__card" elevation="10">
            <font-awesome-layers class="fa-3x service-item__icon">
              <font-awesome-icon class="service-item__icon--background" :icon="['fas', 'circle']" />
              <font-awesome-icon
                class="service-item__icon--plus"
                :icon="['fas', 'plus']"
                transform="shrink-6"
              />
            </font-awesome-layers>

            <v-img
              src="https://cdn.vuetifyjs.com/images/cards/docks.jpg"
              class="service-item__background white--text"
              gradient="120deg, rgba(0,0,0,0.88) 0%, rgba(0,0,0,0.6) 45%, rgba(0,0,0,0.4) 60%,  rgba(0,0,0,0) 100%"
            >
              <h2 class="service-item__title">{{ service.title }}</h2>
              <p class="light">{{ service.serviceText }}</p>
            </v-img>
            <div class="service-item__big-description">
              <div class="service-item__big-description--wrap">
                <p>{{ service.bigDescription }}</p>
              </div>
            </div>
          </v-card>
        </v-col>
      </div>
    </div>
  </div>
</template>
<script>
import { library } from "@fortawesome/fontawesome-svg-core";
import { faPlus, faCircle } from "@fortawesome/free-solid-svg-icons";
import {
  FontAwesomeIcon,
  FontAwesomeLayers
} from "@fortawesome/vue-fontawesome";

library.add(faPlus, faCircle);
import JQuery from "jquery";
var $ = JQuery;
export default {
  name: "services",
  components: {
    FontAwesomeIcon,
    FontAwesomeLayers
  },
  data: () => ({
    links: false,
    hover: false,
    services: [
      {
        id: 1,
        title: "Service 1",
        serviceText: "This is a small description of the service.",
        bigDescription:
          "Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo."
      },
      {
        id: 2,
        title: "Service 2",
        serviceText: "This is a small description of the service.",
        bigDescription:
          "Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo."
      },
      {
        id: 3,
        title: "Service 3",
        serviceText: "This is a small description of the service.",
        bigDescription:
          "Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo."
      },
      {
        id: 4,
        title: "Service 4",
        serviceText: "This is a small description of the service.",
        bigDescription:
          "Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo."
      },
      {
        id: 5,
        title: "Service 5",
        serviceText: "This is a small description of the service.",
        bigDescription:
          "Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo."
      },
      {
        id: 6,
        title: "Service 6",
        serviceText: "This is a small description of the service.",
        bigDescription:
          "Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo."
      }
    ],
    oddServices: [],
    evenServices: [],
    loaded: false,
    enableScroll: false,
    destroyComponent: true
  }),
  mounted() {
    //const intro = this.$refs.intro;
    //const timeline = new TimelineLite();
    //gsap.to(".services-page__intro--title", { rotation: 27 });
    // var card = document.querySelectorAll(".service-item__big-description");
    // var log = document.getElementById("log");
    // card.on("animationstart webkitAnimationStart", function() {
    //   log.empty();
    //   log.append("<p>Animation has started</p>");
    // });
    // card.on("animationiteration webkitAnimationIteration", function() {
    //   log.append("<p>An iteration fired</p>");
    // });
  },
  methods: {
    mouseEnter: function(event) {
      var eventClass = event.srcElement;
      eventClass.classList.add("service-item--hover");
      if (eventClass.className.indexOf("sevice-item--hover")) {
        this.hover = true;
      }
    },
    mouseLeave: function(event) {
      var eventClass = event.srcElement;
      eventClass.classList.remove("service-item--hover");
      this.hover = false;
    },
    getServices() {
      this.loaded = true;

      for (var i = 0; i < this.services.length; i++) {
        if (this.services[i].id % 2 === 0) {
          this.evenServices.push(this.services[i]);
        } else {
          this.oddServices.push(this.services[i]);
        }
      }
    },
    serviceSelect: function(event) {
      //const tl = gsap.timeline();
      if (event) {
        //tl.to(".services-page__intro--title", { duration: 1, rotation: 90 });
        $(event.target)
          .closest(".service-item")
          .toggleClass("service-item--active");
      }

      /* eslint-disable no-alert, no-console */
      console.log(
        $(".service-item--active .service-item__big-description").height()
      );
      /* eslint-enable no-alert, no-console */
    },
    titleSelect: function(event) {
      if (event) {
        $(event.target)
          .closest(".services-page__intro")
          .toggleClass("services-page__intro--active");
      }
    },
    destroyComp: function() {
      this.destroyComponent = false;
    }
  },
  created() {
    this.getServices();
  },
  destroyed() {
    this.destroyComp();
  }
};
</script>
<style lang="scss">
.layerd-icons {
  position: relative;
  height: 50px;
  width: 50px;
}
.services-page {
  padding-top: 5rem;
  width: 100%;

  .page-sections {
    display: none;
  }

  @include respond(mobileSmall) {
    padding-top: 10rem;
  }

  .row {
    max-width: 150rem;
    margin: auto;
    display: grid;
    column-gap: 2rem;
    grid-template-columns: 100%;
    @include respond(mobileLarge) {
      grid-template-columns: 50% 50%;
    }
  }

  &__intro {
    width: 100%;

    grid-column: span 1;

    @include respond(mobileLarge) {
      grid-column: span 2;
    }

    &--title {
      text-align: center;
      line-height: 11rem;
    }
  }

  .service-item {
    display: grid;
    position: relative;
    max-width: 100%;
    //flex: 1 0 100%;

    @include respond(mobileSmall) {
      padding: 12px 20px;
    }

    @include respond(desktopSmall) {
      padding: 12px;
    }
    cursor: pointer;
    &__title {
      letter-spacing: 0.3rem;
    }
    &__background {
      height: 100%;
      align-items: flex-end;
      padding: 5rem 2rem;
      z-index: 3;

      @include respond(mobileSmall) {
        height: 300px;
      }
    }

    &__icon {
      justify-self: end;
      right: -15px;
      top: -16px;
      z-index: 4;
      position: absolute;
      backface-visibility: hidden;
      &--background {
        color: $primary-dark;
      }

      &--plus {
        // transform: rotate(0) scale3d(1, 1, 1) translateZ(0);
        -webkit-transition: all 0.3s cubic-bezier(0.25, 0.8, 0.5, 1);
        transition: all 0.3s cubic-bezier(0.25, 0.8, 0.5, 1);
        color: $color-white;
        backface-visibility: hidden;
        -webkit-transform: scale(0.7);
        transform: scale(0.7);
        -webkit-font-smoothing: subpixel-antialiased;
      }
    }

    &__card {
      //height: 300px;
      display: grid;
      align-self: start;
      position: relative;
    }

    &__big-description {
      @include respond(mobileSmall) {
        opacity: 0;
        height: 0;
        -webkit-transform: translateY(-154px);
        transform: translateY(-154px);
        -webkit-transition: all 0.3s ease-in-out;
        transition: all 0.3s ease-in-out;
      }

      &--wrap {
        padding: 12px 20px 20px 20px;
        transform: translateY(0);
        background-color: transparent;
        margin-bottom: 0;
      }
    }

    &--hover {
      .service-item__icon--plus {
        -webkit-transform: scale(1);
        backface-visibility: hidden;
        -webkit-font-smoothing: subpixel-antialiased;
      }
    }

    &--active {
      .service-item__big-description {
        display: flex;
        @include respond(mobileSmall) {
          height: 187px;
          opacity: 1;
          -webkit-transform: translateY(0);
          transform: translateY(0);
          -webkit-transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
          transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
          animation: none;
        }
        @include respond(mobileSmallLand) {
          height: 100px;
          opacity: 1;
          -webkit-transform: translateY(0);
          transform: translateY(0);
          -webkit-transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
          transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
          animation: none;
        }
        @include respond(mobileLarge) {
          height: 190px;
          opacity: 1;
          -webkit-transform: translateY(0);
          -webkit-transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
          transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
          animation: none;
        }
        @include respond(tabletSmall) {
          height: 200px;
          opacity: 1;
          -webkit-transform: translateY(0);
          transform: translateY(0);
          -webkit-transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
          transition: all 0.6s cubic-bezier(0.67, 0.01, 0.52, 1.4);
          animation: none;
        }
        @include respond(tabletLarge) {
        }
      }
      .service-item {
        &__card {
          height: auto;
        }

        &__icon--plus {
          transform: rotate(225deg);
          transition: all 0.3s cubic-bezier(0.25, 0.8, 0.5, 1);
        }
      }
    }
  }
}

@-webkit-keyframes -webkit-scale {
  0% {
    -webkit-transform: scale(1);
    transform: scale(1);
    height: 50px;
  }
  50% {
    -webkit-transform: scale(1.2);
    transform: scale(1.2);
    height: 70px;
  }
  100% {
    -webkit-transform: scale(1.5);
    transform: scale(1.5);
    height: 100px;
  }
}
@keyframes scale {
  0% {
    -webkit-transform: scale(1);
    transform: scale(1);
    height: 50px;
  }
  50% {
    -webkit-transform: scale(1.2);
    transform: scale(1.2);
    height: 70px;
  }
  100% {
    -webkit-transform: scale(1.5);
    transform: scale(1.2);
    height: 100px;
  }
}

@-webkit-keyframes slideDownDesktop {
  0% {
    opacity: 0;
    -webkit-transform: translateY(-154px);
    //transform: translateY(-154px);
    height: 0px;
  }
  50% {
    opacity: 1;
    -webkit-transform: translateY(20px);
    //transform: translateY(20px);
    height: 200px;
  }
  100% {
    opacity: 1;
    -webkit-transform: translateY(0);
    //transform: translateY(0);
    height: 190px;
  }
}
@keyframes slideDownDesktop {
  0% {
    opacity: 0;
    transform: translateY(-154px);
    height: 0px;
  }
  50% {
    opacity: 1;
    transform: translateY(20px);
    height: 200px;
  }
  100% {
    opacity: 1;
    transform: translateY(0);
    height: 190px;
  }
}
@-webkit-keyframes slideDownTablet {
  0% {
    opacity: 0;
    -webkit-transform: translate(-154px);
    transform: translateY(-154px);
    height: 0px;
  }
  50% {
    opacity: 1;
    -webkit-transform: translateY(20px);
    transform: translateY(20px);
    height: 210px;
  }
  100% {
    opacity: 1;
    -webkit-transform: translateY(0);
    transform: translateY(0);
    height: 200px;
  }
}
@keyframes slideDownTablet {
  0% {
    opacity: 0;
    -webkit-transform: translate(-154px);
    transform: translateY(-154px);
    height: 0px;
  }
  50% {
    opacity: 1;
    -webkit-transform: translateY(20px);
    transform: translateY(20px);
    height: 210px;
  }
  100% {
    opacity: 1;
    -webkit-transform: translateY(0);
    transform: translateY(0);
    height: 200px;
  }
}
@-webkit-keyframes slideDownMobile {
  0% {
    opacity: 0;
    -webkit-transform: translateY(-154px);
    transform: translateY(-154px);
    height: 0px;
  }
  50% {
    opacity: 1;
    -webkit-transform: translateY(20px);
    transform: translateY(20px);
    height: 195px;
  }
  100% {
    opacity: 1;
    -webkit-transform: translateY(0);
    transform: translateY(0);
    height: 187px;
  }
}
@keyframes slideDownMobile {
  0% {
    opacity: 0;
    -webkit-transform: translateY(-154px);
    transform: translateY(-154px);
    height: 0px;
  }
  50% {
    opacity: 1;
    -webkit-transform: translateY(20px);
    transform: translateY(20px);
    height: 195px;
  }
  100% {
    opacity: 1;
    -webkit-transform: translateY(0);
    transform: translateY(0);
    height: 187px;
  }
}
@keyframes slideDownLand {
  0% {
    opacity: 0;
    -webkit-transform: translateY(-154px);
    transform: translateY(-154px);
    height: 0px;
  }
  50% {
    opacity: 1;
    -webkit-transform: translateY(20px);
    transform: translateY(20px);
    height: 120px;
  }
  100% {
    opacity: 1;
    -webkit-transform: translateY(0);
    transform: translateY(0);
    height: 100px;
  }
}
</style>
