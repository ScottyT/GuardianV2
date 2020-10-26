<template>
  <div class="slices-wrapper">
    <section class="home-hero-wrapper" v-for="(slice, i) in homeHeroSlice" :key="i">
      <template v-if="slice.slice_type === 'homehero'">
        <home-hero :slice="slice" />
      </template>
    </section>
    <div
      :class="`slices-wrapper__body-wrapper slices-wrapper__body-wrapper--${Object.keys($route.params).length === 0 ? 'home' : 'page'}`">
      <section v-for="(slice, index) in slices" :key="index" class="section"
        :class="slice.slice_type === 'twocolcardlist' ? slice.primary.label : ''">
        <template v-if="slice.slice_type === 'pagehero'">
          <page-hero :slice="slice" />
        </template>
        <template v-else-if="slice.slice_type === 'twocolcardlist'">
          <two-col-cards :slice="slice" />
          <div class="paint-roller" v-if="$vuetify.breakpoint.width > 768 && $route.params.uid === 'services'">
            <img
              src="https://images.prismic.io/guardianrestoration/b04b10b4-c826-4b47-be23-f5fcf8cdb64c_Paint+Roller+for+Our+Services.png?auto=compress,format" />
          </div>
          <!-- <lazy-image source="https://images.prismic.io/guardianrestoration/96c6ffbf-d737-4fcf-bba9-b71cf728d9d3_Roof+Image+Our+Services+Guardian.png?auto=compress,format" imageClass="section__roof-image" v-if="$vuetify.breakpoint.width < 768 && $route.params.uid === 'services'" /> -->
        </template>
        <template v-else-if="slice.slice_type === 'offsetsectiontext'">
          <offset-section-text :slice="slice" />
        </template>
        <template v-else-if="slice.slice_type === 'contactus'">
          <contact-us :slice="slice" />
        </template>
        <template v-else-if="slice.slice_type === 'rollupsection'">
          <roll-up-section :slice="slice" />
        </template>
        <template v-else-if="slice.slice_type === 'beforeandafter'">
          <before-after-section :slice="slice" />
        </template>
        <template v-else-if="slice.slice_type === 'fullwidthimage'">
          <full-width-image :slice="slice" />
        </template>
        <template v-else-if="slice.slice_type === 'singleimage'">
          <single-image :slice="slice" />
        </template>
        <template v-else-if="slice.slice_type === 'listprojects'">
          <list-projects :slice="slice" />
        </template>
        <template v-else-if="slice.slice_type === 'projectcalendar'">
          <div class="gray-bg">
            <project-calendar :slice="slice" />
          </div>
        </template>
        <template v-else-if="slice.slice_type === 'bgvideo'">
          <bg-video :slice="slice" />
        </template>
      </section>
    </div>
    <div class="slices-wrapper__service-buttons">
      <section v-for="slice in serviceButtonSlice" :key="slice.id">
        <template v-if="slice.slice_type === 'servicesiconbuttons'">
          <services-icon-buttons :slice="slice" />
        </template>
      </section>
    </div>
  </div>
</template>
<script>
  export default {
    name: "slices-block",
    props: {
      slices: Array,
    },
    computed: {
      homeHeroSlice() {
        var slicesArr = this.slices
        const heroSliceIndex = slicesArr.filter((slice) => {
          return slice.slice_type == "homehero"
        })
        return heroSliceIndex
      },
      serviceButtonSlice() {
        var slicesArr = this.slices
        return slicesArr.filter((slice) => {
          return slice.slice_type == "servicesiconbuttons"
        })
      }
      // filteredSlices() { 
      // 	var removeElement = this.slices.shift()
      // 	return this.slices
      // }
    },
    components: {
      TwoColCards: () => import("./slices/TwoColCardList.vue"),
    },
  }
</script>
<style lang="scss">
  .v-list-item__title {
    font-size: 1em !important;
  }

  html {
    @include respond(mobileSmall) {
      font-size: 56.5% !important;
    }

    @include respond(mobileLarge) {
      font-size: 58.5% !important;
    }

    @include respond(tabletLarge) {
      font-size: 62.5% !important;
    }
  }

  .v-application--wrap {
    overflow: hidden;
  }

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

  body {
    font-size: 18px;
  }

  h1 {
    font-family: $heading-font;
    letter-spacing: 1px;

    @include respond(desktopSmall) {
      font-size: 4em;
    }

    &.extra-large {
      font-size: 2.5em;

      @include respond(mobileLarge) {
        font-size: 3.5em;
      }

      @include respond(tabletLarge) {
        font-size: 5em;
      }

      @include respond(desktopSmall) {
        font-size: 6em;
      }
    }

    &.red-text {
      color:$primary-dark;
    }
    &.small {
      font-size:2em;
      @include respond(tabletLarge) {
        font-size:3em;
      }
    }
  }

  h2 {
    font-size: 1.7em;
    line-height: 1.3;
    font-family: $heading-font;

    @include respond(mobileLarge) {
      font-size: 2em;
    }

    @include respond(tabletLarge) {
      font-size: 2.5em;
    }

    @include respond(desktopSmall) {
      font-size: 3em;
    }
  }

  h3 {
    font-size: 2.027em;
  }

  h4 {
    font-size: 1.424em;
  }

  ul {
    list-style: none;
  }

  .button {
    padding: 4px 20px;
    display: inline-block;

    &--red {
      background: rgb(251, 132, 132);
      background: linear-gradient(90deg, rgba(251, 132, 132, 1) 0%, rgba(211, 83, 83, 1) 50%);
      transition: box-shadow .3s ease-in;
      color: $color-white;

      &:hover {
        box-shadow: inset 0 0 0 4px $color-black;
        transition: box-shadow .3s ease-in;
      }
    }

    &__call-button {
      padding: 7px 15px 7px 55px;
      border: 2px solid $primary;
      font-weight: 700;
      font-size: .9em;
      background-color:$color-white;
      font-family: $heading-font;
      letter-spacing: 3px;
      position: relative;
      color:$color-black!important;

      @include respond(mobileLarge) {
        font-size: 1.2em;
      }

      &::after {
        content: '';
        position: absolute;
        background: url('https://images.prismic.io/guardianrestoration/9e332084-cc94-4611-abac-010984e892f4_CALL+NOW+button+3+waves.png?auto=compress,format');
        background-size: contain;
        background-repeat: no-repeat;
        width: 51px;
        height: 100%;
        right: -60px;
        top: 0;
      }

      &::before {
        content: '';
        position: absolute;
        background: url('https://images.prismic.io/guardianrestoration/d6565558-4e9f-431a-ac43-519d90904eae_Phone+for+Call+Now+button+Guardian.png?auto=compress,format');
        background-size: contain;
        background-repeat: no-repeat;
        width: 36px;
        height: 36px;
        left: 11px;
        top: 5px;

        @include respond(mobileLargeMax) {
          height: 25px;
        }
      }
    }
  }

  p {
    @include respond(mobileLargeMax) {
      font-size: .9em;
    }
  }

  a {
    color: inherit;
    text-decoration: none;

    &.color-red {
      color: $primary-dark;
    }

    &.button {
      //	display:inline-block;
      padding: 4px 14px;
      max-width: 100px;
      width: 100%;
      color: $color-white;

      &--red {
        background: rgb(251, 132, 132);
        background: linear-gradient(90deg, rgba(251, 132, 132, 1) 0%, rgba(211, 83, 83, 1) 50%);
        transition: box-shadow .3s ease-in;

        &:hover {
          box-shadow: inset 0 0 0 4px $color-black;
          transition: box-shadow .3s ease-in;
        }
      }

      &--expand {
        position: absolute;
        top: -47px;
        transform: translateX(-50%);
        left: 50%;
        text-align: center;
        text-transform: uppercase;
        max-width: 150px;
        z-index: 3;

        @include respond(mobileLarge) {
          top: -67px;
        }
      }

      &--services {
        color: $color-white;
        position: absolute;
        left: 50%;
        transform: translateX(-50%) translateY(-5px);
        font-size: 1.2em;
        text-transform: uppercase;
        font-family: $heading-font;
        z-index: -1;
        bottom: 23px;
        opacity: 0;
        transition: opacity .5s ease-in, transform .5s ease-in;
      }
    }
  }

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  .slices-wrapper {
    position: relative;

    &__body-wrapper {
      background-image: url('https://images.prismic.io/guardianrestoration/9d94a63c-fe7f-461b-bc1f-c5b446dccd79_Triangles+BG+for+Guardian+home.png?auto=compress,format');
      background-repeat: repeat-y;
			background-size:cover;
			background-position-x:51px;
			@include respond(mobileLarge) {
				background-size: contain;
			}

      &--home {
        background-image: url('https://images.prismic.io/guardianrestoration/9d94a63c-fe7f-461b-bc1f-c5b446dccd79_Triangles+BG+for+Guardian+home.png?auto=compress,format');
        background-repeat: repeat;
        padding-top: 30px;			
				@include respond(mobileLarge) {
					background-size: contain;
				}
      }

      &--page {
        .before-after-block {
          width: 100%;

          @include respond(mobileSmallPort) {
            max-width: 87%;
            height: 416px;
          }

          @include respond(mobileLarge) {
            max-width: 700px;
          }

          @include respond(tabletLarge) {
            max-width: 940px;
            height: 531px;
          }
        }
      }

      @include respond(mobileLarge) {
        background-position-y: -305px;
        background-size: cover;
      }
    }
  }

  .section {
    position: relative;
    height: auto;

    &__roof-image {
      @include respond(mobileLargeMax) {
        height: 276px;
      }
    }

    &__services-cards {
      background-image: url('https://images.prismic.io/guardianrestoration/69818d1d-e79a-496e-9a54-8ca3de0c10cc_Concrete+Wall+for+Our+Services+Guardian+3.3+MB.jpg?auto=compress,format');
      background-size: cover;
      position: relative;
      z-index: 2;
      box-shadow: 0px 40px 12px 0px rgba(6, 6, 6, .52);

      @include respond(mobileLarge) {
        height: 1713px;
        box-shadow: 0px 60px 12px 0px rgba(6, 6, 6, .52);
      }

      &:after {
        content: '';
        position: absolute;
        bottom: 0;
        width: 100%;
				height: 372px;
				background-size:cover;
        background-image: url('https://images.prismic.io/guardianrestoration/96c6ffbf-d737-4fcf-bba9-b71cf728d9d3_Roof+Image+Our+Services+Guardian.png?auto=compress,format');

        @include respond(mobileLarge) {
          height: 636px;
        }
      }

      .paint-roller {
        height: 388px;
        position: absolute;
        bottom: 653px;

        @include respond(mobileLargeMax) {
          display: none;
        }

        img {
          object-position: right center;
        }
      }
    }

    &__about-cards {
      background-image: url('https://images.prismic.io/guardianrestoration/69818d1d-e79a-496e-9a54-8ca3de0c10cc_Concrete+Wall+for+Our+Services+Guardian+3.3+MB.jpg?auto=compress,format');

    }
  }

  .form {
    &__input {
      padding: 6px 12px;

      &:not(:last-child) {
        margin-bottom: 5px;
      }
    }
  }

  .modal-backdrop {
    position: fixed;
    top: 0;
    left: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 100%;
    background: rgba($color-black, 0.45);
    z-index: 999;
    transition: opacity 0.4s ease;
  }

  .form {
    &__input {
      padding: 2px 9px;
      display: block;
      border-bottom: 1px solid $primary-dark;

      &--input-group {
        display: grid;
        grid-template-columns: auto 1fr;
        grid-template-rows: 1fr 1fr;
        column-gap: 20px;
      }

      &--error {
        color: red;
      }
    }

    &__label {
      font-weight: 700;

    }
  }

  //generic classes
  .text-center {
    text-align:center;
  }

  .block-heading {
    position: relative;
  }

  .gray-bg {
    background-color: $color-light-grey;
  }

  // transition stuff
  .fadeIn-enter {
    opacity: 0;
  }

  .fadeIn-leave-active {
    opacity: 0;
    transition: all 0.4s ease-in;
  }

  .fadeIn-enter .modal,
  .fadeIn-leave-active.modal {
    transform: scale(1.1);
  }
</style>