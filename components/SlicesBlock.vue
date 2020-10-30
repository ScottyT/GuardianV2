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
    font-size:1.2em;
    @include respond(mobileLarge) {
      font-size:1.5em;
    }
    @include respond(tabletLarge) {
      font-size: 2.027em;
    }
  }

  h4 {
    font-size: 1.424em;
  }

  ul {
    list-style: none;
  }

  input {
    &:focus {
      outline:none!important;
    }
  }
  textarea {
    &:focus {
      outline:none!important;
    }
  }
  button {
    outline:none!important;
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

    &--submit {
      background-color: #fdc2c2;
      font-size:.9em;
      font-weight:700;
      font-family:'Oswald', sans-serif;
      padding:5px 22px;
      transition:all .3s ease-in;
      @include respond(mobileSmallPort) {
        font-size:1.2em;
        padding:5px 30px;
      }
      @include respond(tabletLarge) {
        padding:10px 50px;
        font-size:1.4em;
      }
      &.pushed {
       // background-color:$primary;
      }
      &.submitting {
        background-color:$primary;
        transition:all .3s ease-in;
        color:$color-white;
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

  .wing {
    position:absolute;
    max-width:120px;
    object-fit:contain;
    opacity:0;
    transition:all .8s cubic-bezier(0.33, 0.01, 0.13, 1.46);
    z-index:-1;
    top:0;
    @include respond(mobileSmallPort) {
      max-width:150px;
    }
    @include respond(mobileLarge) {
      max-width:209px;
    }
    &--right {
      transform:rotateY(180deg);
      right:0;
      
      &.visible {
        right:-115px;
        opacity:1;
        z-index:1;
        @include respond(mobileSmallPort) {
          right:-150px;
        }
        @include respond(mobileLarge) {
          right:-234px;
        }
      }
    }
    &--left {
      left:0px;
      &.visible {
        opacity:1;
        left:-115px;
        z-index:1;
        @include respond(mobileSmallPort) {
          left:-150px;
        }
        @include respond(mobileLarge) {
          left:-232px;
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

    &--contact {
      .button {
        box-shadow:-3px 4px 9px 0px rgba($color-black, .7);
        transition:all .2s ease-in-out;
        &.pushed {
          box-shadow:-1px 1px 3px 0px rgba($color-black, .7);
        }
      }
      .button-bg {
        background-color:$color-white;
        border-radius:10px;
        padding:10px;
        margin:auto;
        position:relative;
        transition:all .3s ease-in;
      }
    }

    &__rounded-input-wrapper {
        display:flex;
        width:100%;
        position:relative;
        flex-direction:column;
        height:74px;
        align-items:center;
        margin-bottom:30px;
        border-radius:20px;
        background-color:$color-white;
        padding:10px;
        
        @include respond(mobileLarge) {
          height:84px;
          flex-direction:row;
        }

        & > div {
          width:100%;
          height:100%;
          display:flex;
          align-items:center;
        }

        &--textarea {
          height:300px;
          align-items:flex-start;
        }
        .form__input--error {
          padding-left:10px;
          @include respond(mobileLargeMax) {
            font-size:.9em;
          }
        }
      }

    &__rounded-input {
      border-radius:20px;
      box-shadow:-3px 4px 5px 0px rgba($color-black, .5);
      border:1px solid $color-grey;
      padding:10px 15px 10px 46px;
      height:100%;
      width:100%;      
      transition: background-color 0.3s, color 0.3s ease-in-out, box-shadow .2s ease-out;

      &.is-true {
        background-color: #0AE569;
        color: #045929;
        box-shadow:-1px 1px 5px 0px rgba($color-black, .5);
      }
      &.is-false {
        background-color: #FFA4A2;
        color: #EB0600;
      }      

      &:focus {
        background-color:#8CFFFE;
        border:2px solid #8CFFFE;
        box-shadow:-1px 1px 5px 0px rgba($color-black, .5);
        color:#000;
      }

      &--message {
        resize:none!important;
      }

      &::placeholder {
        color:black;
        font-weight:700;
        font-family:'Oswald', sans-serif;
        font-size:1em;
        @include respond(mobileLarge) {
          font-size:1.2em;
        }
      }
      &:-ms-input-placeholder {
        color:black;
        font-weight:700;
        font-family:'Oswald', sans-serif;
        font-size:1em;
        @include respond(mobileLarge) {
          font-size:1.2em;
        }
      }
      &::-ms-input-placehodler {
        color:black;
        font-weight:700;
        font-family:'Oswald', sans-serif;
        font-size:1em;
        @include respond(mobileLarge) {
          font-size:1.2em;
        }
      }
    }
    &__icon {
      width:36px;
      height:37px;
      position:absolute;
      left:18px;
      @include respond(mobileLargeMax) {
        top:12px;
      }

      &--textarea-icon {
        top:20px;
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