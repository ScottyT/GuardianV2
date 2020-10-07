<template>
  <div class="offset-section-text">
    <div class="offset-section-text__row before-after-row">
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
        <div class="scroller">
          <svg class="scroller__thumb" xmlns="http://www.w3.org/2000/svg" width="100" height="100" viewBox="0 0 100 100"><polygon points="0 50 37 68 37 32 0 50" style="fill:#fff"/><polygon points="100 50 64 32 64 68 100 50" style="fill:#fff"/></svg>
        </div>
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
    this.$nextTick(() => {
      const imageWrapper = document.querySelector('.before-after-block')
      const beforeLayer = imageWrapper.querySelector('.before-after-block__image--before')
      const afterLayer = imageWrapper.querySelector('.before-after-block__image--after')
      const scroller = document.querySelector('.scroller');
      var delta = 0
      // image.addEventListener('mousemove', (e) => {
      //   console.log(e.clientX)
      //   delta = (e.clientX - image.offsetWidth / 2)
      //   beforeLayer.style.width = delta + 'px'
      // })

      let active = false;
      scroller.addEventListener('mousedown', function() {
        active = true;
        scroller.classList.add('scrolling')
      });
      document.body.addEventListener('mouseup', function() {
        active = false;
        scroller.classList.remove('scrolling')
      });
      document.body.addEventListener('mouseleave', function() {
        active = false;
        document.querySelector('.scroller').classList.remove('scrolling')
      });

      document.body.addEventListener('mousemove', function(e) {
        if (!active) return;
        let x = e.pageX;
        x -=
        imageWrapper.getBoundingClientRect().left;
        scrollIt(x)
      });
      function scrollIt(x) {
        let transform = Math.max(0, (Math.min(x, imageWrapper.offsetWidth)));
        beforeLayer.style.width = transform+"px";
        scroller.style.left = transform-25+"px";
      }
      scrollIt(250)
      scroller.addEventListener('touchstart', function() {
        active = true;
        scroller.classList.add('scrolling');
      });
      document.body.addEventListener('touchend', function() {
        active = false;
        scroller.classList.remove('scrolling');
      });
      document.body.addEventListener('touchcancel', function() {
        active = false;
        scroller.classList.remove('scrolling')
      });
    })
  }
}
</script>
<style lang="scss">
.scroller{
  width: 50px;
  height:50px;
  position: absolute;
  left:100px;
  top:50%;
  transform:translateY(-50%);
  border-radius:50%;
  border: 5px solid $color-white;
  background-color: transparent;
  opacity:0.9;
  pointer-events:auto;
  z-index:2;
  cursor: pointer;

  &:hover {
    opacity:1
  }
  &__thumb {
    width:100%;
    height:100%;
    padding:5px;
  }

  &:before, &:after {
    content:" ";
    background:$color-white;
    display: block;
    width: 7px;
    height: 9999px;
    position: absolute;
    left: 50%;
    margin-left: -3.5px;
    z-index: 30;
    transition:0.1s;
  }
  &:before {
    top:100%;
  }
  &:after {
    bottom:100%;
  }
}
.scrolling{
  pointer-events:none;
  opacity:1;
  // z-index: 1;
}

.before-after-row {
  max-width:1200px;
  margin:auto;
}
.before-after-block {
  position:relative;
  overflow:hidden;
  height:400px;

  &__text {
    position:absolute;
    top:50%;
    transform:translateX(-50%);
    z-index:3;
    background-color:rgba($color-black, .5);
    padding:5px;
    color:$color-white;
    
    &--before {
      left:40px;
      transition:left .3s ease-in;
    }
    &--after {
      right:0;
      transition:right .3s ease-in;
    }
  }

  &:hover {
    .before-after-block__text--before {
      left:-40px;
      transition:left .3s ease-in;
    }
    .before-after-block__text--after {
      right:-115px;
      transition:right .3s ease-in;
    }
  }

  &__image {
    position:absolute;
    //max-width:650px;
    overflow:hidden;
    height:100%;
    pointer-events:none;
    top:0;
    left:0;

    &--before {
      width:100%;
      height:100%;
      z-index:1;
      //margin-left:-50px;
      @include respond(mobileSmallPort) {
        //width:50%;
      }
    }
    &--after {
      width:100%;
      height:100%;
    }

    img {
      width:auto;
      height:100%;
    }
  }
  &__before {
    margin-left:50px;
    height:100%;
  }
}
</style>