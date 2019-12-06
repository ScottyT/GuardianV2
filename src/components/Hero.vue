<template>
  <div style="position:relative;" id="hero">
    <v-img class="hero" max-width="100%" :lazy-src="`${hero}`" :src="hero" :gradient="gradient">
      <div class="hero__wrapper">
        <h1 class="extra-large white--text mb-5 hero__wrapper--heading">{{ ctaHeading }}</h1>
        <p class="white--text font-weight-light hero__wrapper--subheading">{{ ctaSubheading }}</p>
      </div>
    </v-img>
  </div>
</template>
<script>
export default {
  name: "Hero",
  props: {
    hero: String,
    ctaHeading: String,
    ctaSubheading: String,
    gradient: String
  },
  mounted() {}
};
</script>
<style lang="scss">
.home {
  .hero {
    @include respond(mobileSmall) {
      height: calc(100vh - 150px);
    }

    @include respond(mobileSmallLand) {
      height: calc(100vh - 100px);
    }

    @include respond(iPhoneX) {
      height: calc(100vh - 100px);
    }
    @include respond(GalaxyS10Plus) {
      height: calc(100vh - 100px);
    }
  }
}
.hero {
  &__wrapper {
    max-width: 800px;
    width: 79%;
    position: absolute;

    @include respond(mobileSmall) {
      top: 55%;
      transform: translateY(-50%);
    }
    @include respond(desktopSmall) {
      max-width: auto;
    }

    @media (min-width: 768px) {
      @include absCentered(50%, 50%);
    }

    &--subheading {
      max-width: 100%;

      @include respond(mobileLarge) {
        max-width: 70%;
        font-size: 1.8rem;
      }
    }

    @media only screen and (max-width: 767px) {
      max-width: 100%;
      width: auto;
      padding: 0 5rem;
    }
  }
}
</style>
