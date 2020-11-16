<template>
  <div class="services-featured" :style="`background-image:url(${slice.primary.bg_image.url}`">
    <div :class="`services-featured__col services-featured__col--${i}`" :data-col="i" v-for="(column, i) in slice.items"
      :key="i">
      <span style="text-align:center;" v-if="$vuetify.breakpoint.width < 768">
        <div class="services-featured__heading">{{$prismic.asText(column.col_title)}}</div>
        <div class="services-featured__subheading">{{$prismic.asText(column.col_subtitle)}}</div>
      </span>

      <div :class="`${isMobile ? 'target--mobile' : 'target'}`">
        <span :class="`services-featured__col-bg`"><span class="sr-only">Bg</span></span>
        <div class="services-featured__heading" v-if="$vuetify.breakpoint.width > 768">
          {{$prismic.asText(column.col_title)}}</div>
        <div class="services-featured__subheading" v-if="$vuetify.breakpoint.width > 768">
          {{$prismic.asText(column.col_subtitle)}}</div>
        <p class="services-featured__text">{{$prismic.asText(column.col_text)}}</p>
      </div>
    </div>
  </div>
</template>
<script>
  import anime from 'animejs';
  import $debounce from 'lodash.debounce';
  import debounce from 'lodash.debounce';
  export default {
    name: "RollUpSection",
    props: ['slice'],
    data: () => ({
      isMobile: false
    }),
    beforeDestroy() {
      if (typeof window === 'undefined') return
      window.removeEventListener('resize', this.onResize, {
        passive: true
      })
    },
    mounted() {
      this.onResize()
      window.addEventListener('resize', this.onResize, {
        passive: true
      })

      function hoverAnimOn(target) {
        var tl = anime.timeline({
          targets: `.${target} .target`,
          translateY: [542, 0],
          easing: 'easeOutElastic(1, 1)',
          duration: 1000,
          begin: function (anim) {
            document.querySelector(`.${target} .services-featured__heading`).classList.add('glow')
            document.querySelector(`.${target} .services-featured__subheading`).classList.add('glow')
            document.querySelector(`.${target} .services-featured__text`).classList.add('glow')
          },
          complete: function (anim) {
            document.querySelector(`.${target} .services-featured__heading`).classList.remove('glow')
            document.querySelector(`.${target} .services-featured__subheading`).classList.remove('glow')
            document.querySelector(`.${target} .services-featured__text`).classList.remove('glow')
          }
        })
        tl.add({
          targets: `.${target} .services-featured__col-bg`,
          opacity: .9
        })
      }

      function hoverAnimOff(target) {
        var tl = anime.timeline({
          targets: `.${target} .target`,
          translateY: [0, 542],
          easing: 'spring(1, 90, 15, 0)',
          duration: 4000
        })
        tl.add({
          targets: `.${target} .services-featured__col-bg`,
          opacity: 0
        })
      }
      var servicesBox = document.querySelector('.services-featured');
      let currentElm = null;

      servicesBox.onmouseover = function (event) {
        if (currentElm) return;

        let target = event.target.closest('div.services-featured__col');
        if (!target) return;
        if (!servicesBox.contains(target)) return;
        currentElm = target
        hoverAnimOn(currentElm.classList[1])
      };
      servicesBox.onmouseout = function (event) {
        if (!currentElm) return;
        let relatedTarget = event.relatedTarget;
        while (relatedTarget) {
          if (relatedTarget == currentElm) return;
          relatedTarget = relatedTarget.parentNode;
        }
        hoverAnimOff(currentElm.classList[1])
        currentElm = null;
      };

    },
    methods: {
      onResize: $debounce(function () {
        this.isMobile = window.innerWidth < 768
      }, 50)
    }
  }
</script>
<style lang="scss">
  .glow {
    color: $color-white;
    -webkit-animation: glow .5s ease-in-out 2 alternate;
    -moz-animation: glow .5s ease-in-out 2 alternate;
    animation: glow .5s ease-in-out 2 alternate;
  }

  @-webkit-keyframes glow {
    from {
      text-shadow: 0 0 10px rgba(255, 255, 255, 0);
    }

    to {
      text-shadow: 0 0 10px #d4cdcd, 0 0 20px #ffffff, 0 0 30px #ffffff, 0 0 40px #000000, 0 0 20px #ffffff, 0 0 60px #9554b300, -1px 11px 70px #e6e2e8;
    }
  }

  .services-featured {
    height:auto;
    width: 100%;
    max-width: 1600px;
    margin: 0 auto 14px;
    background-size: cover;
    color: $color-white;
    // display:grid;
    // grid-template-columns: repeat(4, 1fr);
    overflow: hidden;
    display: block;
    justify-content: space-between;
    flex-direction: column;

    @include respond(mobileLarge) {
      flex-direction: row;
      display: flex;
      justify-content: space-around;
      height: 648px;
      margin: 0 auto 50px;
    }

    &__heading {
      font-family: $heading-font;
      line-height: 1.1;
      z-index: 1;
      font-size: 2em;
      text-transform: uppercase;

      @include respond(desktopSmall) {
        font-size: 2.7em;
      }
    }

    &__subheading {
      font-size: 1.5em;
      font-family: $heading-font;
      z-index: 1;

    }

    &__text {
      margin-top: 10px;
      z-index: 1;
      font-size: .9em;

      @include respond(mobileLargeMax) {
        //display: none;
      }

      @include respond(desktopSmall) {
        font-size: 1em;
      }
    }

    &__col {
      position: relative;

      &>span {
        z-index: 1;
        position: relative;
        display: block;
      }

      &:hover {
        .target--mobile {
          height:250px;
          transition: all .3s ease-in;
          @include respond(mobileSmall) {
            height:200px;
          }
        }
        @include respond(mobileLargeMax) {
          .services-featured__text {
            display: block;
          }

          .services-featured__col-bg {
            opacity: .8;
          }
        }
      }

      @include respond(mobileLarge) {
        width: 24%;
        height: 105%;
      }

      .target {
        overflow: hidden;
        border-left: 11px solid $primary-dark;
        height: 0;
        padding: 10px;
        display: flex;
        flex-direction: column;
        align-items: center;

        &--mobile {
          transition: all .3s ease-in;
          overflow: hidden;
          border-left: 11px solid $primary-dark;
          height: 0;
          padding: 10px;
          display: flex;
          flex-direction: column;
          align-items: center;
        }

        @include respond(mobileLarge) {
          height: 100%;
          overflow: unset;
          transform: translateY(542px);
        }
      }
    }

    &__col-bg {
      z-index: -1;
      position: absolute;
      width: 100%;
      height: 100%;
      top: 0;
      left: 0;
      opacity: 0;
      background-color: rgba($color-grey, 1);

      @include respond(mobileLargeMax) {
        z-index: 0;
        right: 0;
        left: unset;
        width: 98%;
        //transition: opacity .3s ease-in;
      }
    }
  }
</style>