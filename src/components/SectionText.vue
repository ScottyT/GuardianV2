<template>
  <div>
    <v-row class="section-intro">
      <v-col cols="12" md="6" class="offset-col row__column">
        <div class="section-intro__heading row__content--left">
          <h2>{{ heading | allCaps }}</h2>
        </div>

        <div class="row__content row__content--left">
          <p class="subheading" v-show="subheading">{{ subheading }}</p>
          <p v-for="(content, i) in sectionText" v-bind:key="i">{{ content }}</p>
        </div>
      </v-col>
      <div class="section-image-col col-md-6 col-12 no-padding">
        <v-img
          :lazy-src="sectionImage.url"
          :src="sectionImage.url"
          class="section-image-col__img section-image-col__img--right section-image-col__img--big"
        ></v-img>
      </div>
    </v-row>
    <v-row class="section-row">
      <v-col cols="12" md="6" class="section-image-col left no-padding">
        <v-img
          :lazy-src="secondaryImage.url"
          :src="secondaryImage.url"
          class="section-image-col__img section-image-col__img--left section-image-col__img--small"
        ></v-img>
      </v-col>
      <v-col cols="12" md="6" class="justify-center row__column align-center d-flex">
        <div class="row__content row__content--right">
          <p v-for="content in secondaryText" :key="content.id">{{ content }}</p>
        </div>
      </v-col>
    </v-row>
  </div>
</template>
<script>
export default {
  name: "SectionText",
  props: {
    section: Object
  },
  data: () => {
    return {
      heading: "",
      sectionText: [],
      subheading: [],
      sectionImage: [],
      secondaryText: [],
      secondaryImage: []
    };
  },
  methods: {
    getContent() {
      var content = this.section.section_text;
      var sectionText = this.sectionText;
      this.heading = this.section.section_heading[0].text;
      this.sectionImage = this.section.section_image;
      content.map(function(value) {
        sectionText.push(value.text);
      });

      var contentSecond = this.section.secondary_text;
      this.secondaryImage = this.section.secondary_image;
      var secondaryText = this.secondaryText;
      contentSecond.map(function(value) {
        secondaryText.push(value.text);
      });

      this.subheading = this.section.subheading[0].text;
    }
  },
  mounted() {},
  created() {
    this.getContent();
  }
};
</script>
<style lang="scss">
.section {
  position: relative;
  margin: 0 0 4rem 0;
  height: auto;

  @include respond(mobileSmall) {
    margin: 0;
  }

  @include respond(iPhoneX) {
    margin: 0 0 6rem 0;
  }

  @include respond(GalaxyS10Plus) {
    margin: 0 0 6rem 0;
  }

  @include respond(mobileSmallLand) {
    margin: 0;
  }

  @include respond(tabletSmall) {
    margin: 0 0 4rem 0;
  }

  @include respond(tabletLarge) {
    margin: 0 6rem 6rem 0rem;
  }

  @include respond(desktopSmall) {
    margin: 0 7rem 6rem 0;
  }

  &--0 {
    .section-intro {
      &__heading {
        @include respond(iPhoneX) {
          padding-top: 32px;
        }

        @include respond(mobileSmall) {
          padding-top: 0;
        }

        @include respond(tabletLarge) {
        }

        @include respond(desktopSmall) {
          margin-bottom: 10rem;
          padding-top: 53px;
        }
      }

      @include respond(mobileLarge) {
        padding: 0 0 5rem 0;
      }

      .offset-col {
        margin-top: 0;
        padding: 32px 32px 16px;

        @include respond(tabletSmall) {
          margin-top: -7rem;
          padding: 49px 32px 16px;
        }
        @include respond(GalaxyS10Plus) {
          margin-top: 0;
          padding: 50px 32px 16px;
        }
        @include respond(iPhoneX) {
          margin-top: 0;
          padding: 50px 20px 16px 10px;
        }

        @include respond(tabletLarge) {
          margin-top: -7rem;
          //padding: 42px 32px 16px;
          max-width: 100%;
        }

        @include respond(desktopSmall) {
          margin-top: -10rem;
        }
      }

      .section-image-col {
        @include respond(mobileLarge) {
          margin-top: 7rem;
        }
      }
    }

    .section-image-col {
      &__img--right {
        @include respond(mobileSmallLand) {
          // min-height: 230px;
        }
      }
    }
  }

  &--0,
  &--1 {
    &:after {
      content: "";
      display: block;
      position: absolute;
      background-color: $primary-dark;
      height: 76px;
      width: 400px;
      z-index: 4;
      left: 53%;
      transform: translateX(-50%);

      @media only screen and (max-width: $md) {
        display: none;
      }
    }
  }

  &--1,
  &--2 {
    &:before {
      content: "";
      display: block;
      position: absolute;
      background-color: $primary-dark;
      height: 432px;
      width: 100px;
      z-index: 0;
      left: 49%;
      top: 0%;
      transform: translate(-50%, -50%);

      @media only screen and (max-width: $lg) {
        width: 76px;
      }

      @media only screen and (max-width: $md) {
        display: none;
      }
    }
  }

  .section-image-col {
    position: relative;
    height: auto;
    width: 100%;
  }
  .row {
    max-width: 100%;

    @include respond(mobileSmall) {
      margin: 0;
    }

    @include respond(iPhoneX) {
      padding: 0 0 0 0;
      margin-bottom: 5rem;
    }

    @include respond(mobileSmallLand) {
      padding: 0;
      margin: 0;
    }

    &:nth-of-type(2) {
      @include respond(mobileSmall) {
        flex-direction: column-reverse;
      }

      @include respond(mobileLarge) {
        flex-direction: row;
      }
    }

    &__column {
      @media only screen and (max-width: $sm) {
        max-width: 555px;
      }
      @include respond(mobileSmall) {
        margin-top: 0;
        padding: 32px 32px 16px;
        margin: 0 auto;
      }
      @include respond(iPhoneX) {
        padding: 42px 20px 16px 10px;
      }

      @include respond(mobileLarge) {
        display: flex;
        padding: 0 32px 0;
        margin: 0;
      }
      @include respond(tabletSmall) {
        padding: 0 25px 0;
      }

      @include respond(tabletLand) {
        padding: 4px 60px 0 35px;
      }
      @include respond(tabletLarge) {
        padding: 0 40px 0;
      }
      @include respond(desktopSmall) {
        padding: 0 55px 0;
      }
      @include respond(GalaxyS10Plus) {
        padding: 0px 32px 10px;
      }
      @include respond(mobileSmallLand) {
        padding: 24px 32px 10px;
      }
    }

    &__content {
      max-width: 65%;
      z-index: 0;

      @include respond(mobileSmall) {
        padding: 0;
        max-width: 100%;
      }

      &--left {
        @include respond(tabletSmall) {
          padding-right: 20px;
        }

        @include respond(tabletLarge) {
          padding-right: 30px;
          max-width: 555px;
          align-self: flex-end;
        }
      }

      &--right {
        @include respond(tabletSmall) {
          padding-left: 25px;
        }
        @include respond(iPhoneX) {
          padding-left: 25px;
        }
        @include respond(tabletLarge) {
          padding-left: 30px;
          max-width: 555px;
        }
      }
    }
  }
}
.section-row {
  .section-image-col {
    margin: 0;
    padding: 0;
  }
}
.section-intro {
  max-width: 100%;

  @include respond(mobileSmall) {
    padding: 0;
    margin: 0;
  }
  @include respond(iPhoneX) {
    padding: 0;
  }

  @include respond(tabletSmall) {
    padding: 0 0 5rem 0;
  }

  @include respond(tabletLarge) {
    padding: 0 0 3rem 0;
  }
  .section-image-col {
    @include respond(mobileSmall) {
      margin-top: 0;
    }
  }
  .offset-col {
    // padding: 0rem 0;
    background-color: $backgroundColor;
    flex-direction: column;

    @include respond(iPhoneX) {
      margin-top: 0;
      padding: 0px 20px 16px 10px;
    }
    @include respond(mobileLand) {
      margin-top: 0;
    }
    @include respond(mobileSmallLand) {
      padding: 24px 20px 16px 10px;
    }
  }
  &__heading {
    max-width: 65%;
    margin-bottom: 8rem;
    padding-top: 53px;

    @include respond(mobileSmall) {
      padding-top: 0;
      margin-bottom: 4rem;
      max-width: 100%;
    }

    @include respond(mobileSmallLand) {
      margin-bottom: 2.5rem;
      padding-top: 0;
    }
    @include respond(iPhoneX) {
      margin-bottom: 2rem;
    }

    @include respond(tabletLarge) {
      padding-top: 25px;
      padding-right: 20px;
    }
  }
}
.section-image-col {
  z-index: 2;
  &.left {
    justify-content: flex-end;
    display: flex;
    @include respond(tabletSmall) {
      align-items: flex-end;
    }
  }
  .section-intro__content {
    top: 57px;
  }

  &__img {
    //position: absolute;
    @include respond(mobileSmall) {
      min-height: 250px;
    }

    @include respond(iPhoneX) {
      min-height: 290px;
    }
    @include respond(GalaxyS10Plus) {
      min-height: 290px;
    }

    &--big {
      max-width: 625px;

      @include respond(tabletSmall) {
        min-height: 380px;
      }

      @media only screen and (max-width: $lg) {
        max-width: 100%;
      }
    }

    &--small {
      @include respond(mobileSmall) {
        max-width: 100%;
      }

      @include respond(tabletSmall) {
        min-height: 300px;
      }

      @include respond(tabletLarge) {
        max-height: 300px;
      }

      @include respond(desktopSmall) {
        max-width: 555px;
      }
    }

    &--left {
      right: -5px;

      @media only screen and (max-width: $sm) {
        right: 0;
      }

      @include respond(mobileLarge) {
        right: -10px;
      }
    }

    &--right {
      @include respond(mobileSmall) {
        left: 0;
      }

      @include respond(mobileLarge) {
        left: -10px;
      }

      @include respond(tabletLarge) {
        left: -20px;
      }
    }

    @include respond(mobileSmallLand) {
      max-height: 230px;
    }
  }
}
</style>
