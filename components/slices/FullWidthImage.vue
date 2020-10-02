<template>
  <div class="full-width__image" v-lazyload>
    <img :data-url="`${slice.primary.image.url}`" :alt="slice.primary.image.alt" />
    <prismic-rich-text :class="slice.primary.text_class" :field="slice.primary.image_text" />
  </div>
</template>
<script>
import prismicDOM from "prismic-dom";
const Elements = prismicDOM.RichText.Elements;
import LazyLoadDirective from '../../directives/LazyLoadDirective'
export default {
  name:"FullWidthImage",
  props:['slice'],
  directives: {
    lazyload: LazyLoadDirective
  }
}
</script>
<style lang="scss">
.full-width {
  &__image {
    height: 610px;
    width:100%;
    position:relative;
  }
  &__text {
    &--contact-us {
      color:$primary;
      @include absCentered(20%, 20%);
      font-size:2.7em;
      @include respond(mobileLarge) {
        font-size:3em;
      }
      @include respond(tabletLarge) {
        font-size:4em;       
      }
    }
  }
}
</style>