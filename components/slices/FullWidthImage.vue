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
    height:300px;
    width:100%;
    position:relative;

    @include respond(mobileLarge) {
      height: 610px;
    }
    img {     
      object-position:-45px 0px;
      @include respond(mobileLarge) {
        object-position: unset;
      }
    }
  }
  &__text {
    &--contact-us {
      color:$primary;
      @include absCentered(20%, 45%);
      font-size:2em;
      @include respond(mobileLarge) {
        font-size:3em;
        @include absCentered(20%, 20%);
      }
      @include respond(tabletLarge) {
        font-size:4em;       
      }
    }
  }
}
</style>