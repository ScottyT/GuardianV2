<template>
  <div class="offset-section-text">
    <div class="offset-section-text__row before-after-row">
      <div class="offset-section-text__image before-after-block">
        
        <div class="before-after-block__image before-after-block__image--before">
          <div class="before-after-block__before">
            <img :src="slice.primary.before_image.url" :alt="slice.primary.before_image.alt" />
          </div>
        </div>
        <div class="before-after-block__image before-after-block__image--after">
          <img :src="slice.primary.after_image.url" :alt="slice.primary.after_image.alt" />
        </div>
        
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
      const image = document.querySelector('.before-after-block')
      const beforeLayer = image.querySelector('.before-after-block__image--before')
      var delta = 0
      image.addEventListener('mousemove', (e) => {
        console.log(e.clientX)
        delta = (e.clientX - image.offsetWidth / 2)
        beforeLayer.style.width = delta + 'px'
      })
    })
  }
}
</script>
<style lang="scss">
.before-after-row {
  max-width:1200px;
  margin:auto;
}
.before-after-block {
  position:relative;
  overflow:hidden;
  height:400px;

  &__image {
    position:absolute;
    max-width:650px;
    overflow:hidden;
    height:100%;

    &--before {
      width:100%;
      margin-left:-50px;
      z-index:1;
      @include respond(mobileSmallPort) {
        width:50%;
      }
    }

    img {
      width:auto;
    }
  }
  &__before {
    margin-left:50px;
    height:100%;
  }
}
</style>