<template>
  <div class="services-featured" :style="`background-image:url(${slice.primary.bg_image.url}`">
    <div :class="`services-featured__col services-featured__col--${i}`" :data-col="i" v-for="(column, i) in slice.items"
      :key="i">
      <div class="target">
      <span :class="`services-featured__col-bg`"><span class="sr-only">Bg</span></span>
      <div class="services-featured__heading">{{$prismic.asText(column.col_title)}}</div>
      <div class="services-featured__subheading">{{$prismic.asText(column.col_subtitle)}}</div>
      <p class="services-featured__text">{{$prismic.asText(column.col_text)}}</p>
      </div>
    </div>
  </div>
</template>
<script>
  import anime from 'animejs';
  import $debounce from 'lodash.debounce';
  export default {
    name: "RollUpSection",
    props: ['slice'],
    mounted() {
      function hoverAnimOn(target) {
        var tl = anime.timeline({
          targets: `.${target} .target`,
          translateY: [542, 0],
          easing: 'easeOutElastic(1, 1)',
          duration: 1000,
          begin: function(anim) {
            document.querySelector(`.${target} .services-featured__heading`).classList.add('glow')
            document.querySelector(`.${target} .services-featured__subheading`).classList.add('glow')
            document.querySelector(`.${target} .services-featured__text`).classList.add('glow')
          },
          complete: function(anim) {
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
          opacity:0
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
      var serviceCols = document.querySelectorAll('.services-featured__col');
      // serviceCols.forEach((el) => {
      //   console.log(el)
      //   el.addEventListener("mouseover", hoverAnim, false)
      // })
    }
  }
</script>
<style lang="scss">
.glow {
  color:$color-white;
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
    height: 648px;
    width: 100%;
    max-width: 1600px;
    margin: 0 auto 50px;
    background-size: cover;
    color: $color-white;
    // display:grid;
    // grid-template-columns: repeat(4, 1fr);
    overflow: hidden;
    display: flex;
    justify-content: space-around;

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

      @include respond(desktopSmall) {
        font-size: 1em;
      }
    }

    &__col {   
      width: 24%;
      position: relative;
      height:105%;
      
      .target {
        transform: translateY(542px);
        border-left: 11px solid $primary-dark;
        height:100%;
        padding:10px;
        display: flex;
      flex-direction: column;
      align-items: center;
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
    }
  }
</style>