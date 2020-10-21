<template>
  <div :class="`offset-section-text ${slice.primary.cssclass}`">
    <div :class="`offset-section-text__row before-after-row ${slice.primary.label}`">
      <prismic-rich-text :field="slice.primary.heading_copy" v-if="slice.primary.heading_copy.length > 0"
        class="before-after-row__heading-copy" />
      <div class="offset-section-text__image before-after-block">
        <span class="before-after-block__text before-after-block__text--before">Before</span>

        <div class="before-after-block__image before-after-block__image--before">
          <!-- <div class="before-after-block__before">
            
          </div> -->
          <img :src="slice.primary.before_image.url" :alt="slice.primary.before_image.alt" draggable="false" />
        </div>
        <div class="before-after-block__image before-after-block__image--after">
          <img :src="slice.primary.after_image.url" :alt="slice.primary.after_image.alt" draggable="false" />
        </div>
        <div class="scroller-wrapper">
          <div class="scroller-wrapper__tab"></div>
          <div class="scroller">
            <!-- <svg class="scroller__thumb" xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 100 100"><polygon points="0 50 37 68 37 32 0 50" style="fill:#fff"/><polygon points="100 50 64 32 64 68 100 50" style="fill:#fff"/></svg> -->
          </div>
        </div>
        <!-- <div class="scroller" v-if="$route.name === 'index'">
            <svg class="scroller__thumb" xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 100 100"><polygon points="0 50 37 68 37 32 0 50" style="fill:#fff"/><polygon points="100 50 64 32 64 68 100 50" style="fill:#fff"/></svg>
          </div> -->
        <span class="before-after-block__text before-after-block__text--after">After</span>
      </div>
      <div class="offset-section-text__content offset-section-text__content--left">
        <p>{{$prismic.asText(slice.primary.text)}}</p>
      </div>
    </div>
  </div>
</template>
<script>
  export default {
    name: 'BeforeAfterSection',
    props: ['slice'],
    mounted() {
      const imageWrapper = document.querySelector('.before-after-block')
      const beforeLayer = imageWrapper.querySelector('.before-after-block__image--before')
      const afterLayer = imageWrapper.querySelector('.before-after-block__image--after')
      const scroller = document.querySelector('.scroller-wrapper');
      var delta = 0
      // image.addEventListener('mousemove', (e) => {
      //   console.log(e.clientX)
      //   delta = (e.clientX - image.offsetWidth / 2)
      //   beforeLayer.style.width = delta + 'px'
      // })

      let active = false;
      scroller.addEventListener('mousedown', function () {
        active = true;
        scroller.classList.add('scrolling')
      });
      document.body.addEventListener('mouseup', function () {
        active = false;
        scroller.classList.remove('scrolling')
      });
      console.log("hello from before and after component")
      document.body.addEventListener('mouseleave', function () {
        active = false;
        scroller.classList.remove('scrolling')
      });

      document.body.addEventListener('mousemove', function (e) {
        if (!active) return;
        let x = e.pageX;
        x -=
          imageWrapper.getBoundingClientRect().left;
        scrollIt(x)
      });

      function scrollIt(x) {
        let transform = Math.max(0, (Math.min(x, imageWrapper.offsetWidth)));
        beforeLayer.style.width = transform + "px";
        scroller.style.left = transform - 25 + "px";
      }
      scrollIt(250)
      scroller.addEventListener('touchstart', function () {
        active = true;
        scroller.classList.add('scrolling');
        console.log("touch start")
      });
      document.body.addEventListener('touchend', function () {
        active = false;
        scroller.classList.remove('scrolling');
        console.log("touch end")
      });
      document.body.addEventListener('touchmove', function (e) {
        if (!active) return;
        let x = e.touches[0].pageX;
        x -=
          imageWrapper.getBoundingClientRect().left;
        console.log("touch move:", e)
        scrollIt(x)
      });
      document.body.addEventListener('touchcancel', function () {
        active = false;
        scroller.classList.remove('scrolling')
      });
    }
  }
</script>
<style lang="scss">
  .scroller {
    left: 22px;
    border: none;
    width: 0;
    height: 0;

    position: absolute;

    top: 50%;
    transform: translateY(-50%);
    border-radius: 50%;
    background-color: rgba(0, 0, 0, .3);
    opacity: 0.9;
    pointer-events: auto;
    z-index: 2;
    cursor: pointer;

    &:hover {
      opacity: 1
    }

    &__thumb {
      width: 100%;
      height: 100%;
      padding: 5px;
    }

    &:before,
    &:after {
      content: " ";
      background: $primary;
      display: block;
      width: 7px;
      height: 9999px;
      position: absolute;
      left: 50%;
      margin-left: -3.5px;
      z-index: 30;
      transition: 0.1s;
    }

    &:before {
      top: 100%;
    }

    &:after {
      bottom: 100%;
    }
  }

  .scroller-wrapper {
    position: absolute;
    left: 200px;

    &__tab {
      width: 50px;
      height: 150px;
      background-color: $primary;
    }
  }

  .scrolling {
    pointer-events: none;
    opacity: 1;
    // z-index: 1;
  }

  .before-after-row {
    max-width: 1200px;

    &__heading-copy {
      text-align: center;
      width: 100%;
    }

    &__subheading {
      font-size: 1.5em;
      color: $primary;
    }

    &.home {
      .before-after-block {
        @include respond(tabletLargeMax) {
          max-width: 600px;
          width: 100%;
        }
      }
    }

    &.about-us {
      // .scroller-wrapper {
      //   position:absolute;
      //   left:200px;
      //   &__tab {
      //     width:50px;
      //     height:150px;
      //     background-color:$primary;
      //   }
      // }
      flex-direction: column;

      .scroller {
        left: 22px;
        border: none;
        width: 0;
        height: 0;

        &:before,
        &:after {
          background-color: $primary;
        }
      }
    }
  }

  .before-after-block {
    position: relative;
    overflow: hidden;
    height: 314px;

    @include respond(mobileSmallPort) {
      height: 400px;
    }

    &__text {
      position: absolute;
      top: 50%;
      transform: translateX(-50%);
      z-index: 3;
      background-color: rgba($color-black, .85);
      padding: 5px;
      color: $color-white;

      &--before {
        left: 40px;
        transition: left .3s ease-in;
      }

      &--after {
        right: 0;
        transition: right .3s ease-in;
      }
    }

    &:hover {
      .before-after-block__text--before {
        left: -40px;
        transition: left .3s ease-in;
      }

      .before-after-block__text--after {
        right: -115px;
        transition: right .3s ease-in;
      }
    }

    &__image {
      position: absolute;
      //max-width:650px;
      overflow: hidden;
      height: 100%;
      pointer-events: none;
      top: 0;
      left: 0;

      &--before {
        width: 100%;
        height: 100%;
        z-index: 1;

        //margin-left:-50px;
        @include respond(mobileSmallPort) {
          //width:50%;
        }
      }

      &--after {
        width: 100%;
        height: 100%;
      }

      img {
        width: auto;
        height: 100%;
        user-select: none;
        object-position: -75px 0px;

        @include respond(mobileSmallPort) {
          object-position: 0px 0px;
        }
      }
    }

    &__before {
      margin-left: 50px;
      height: 100%;
    }
  }
</style>