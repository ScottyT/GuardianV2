<template>
  <div class="push-slide-wrapper">
    <button class="push-slide-wrapper__button" @click="opened = !opened" role="button">
      PUSH
    </button>
    <div class="push-slide-wrapper__layer">
      <lazy-image :source="image.url" :imageClass="`push-slide-wrapper__image ${opened ? 'slide-open' : '' }`" :alt="image.alt" />
      <!-- <div :class="`push-slide-wrapper__image ${opened ? 'slide-open' : '' }`">
        <img :src="image.url" :alt="image.alt" />
      </div> -->
      <prismic-rich-text :field="hidden" class="push-slide-wrapper__hidden-content" />
    </div>
  </div>
</template>
<script>
export default {
  name: "PushSlide",
  props: ['image', 'element', 'hidden'],
  data() {
    return {
      opened: false
    }
  }
}
</script>
<style lang="scss">
.push-slide-wrapper {
  position:relative;
  max-width:643px;
  height:430px;
  width:100%;
  left:unset;
  right:unset;
  @include respond(mobileLarge) {
    left:20px;
  }
  @include respond(tabletLarge) {
    width:50%;
    //left:-20px;
  }
  &__layer {
    overflow:hidden;
    height:100%;
    position:relative;
    display:flex;
  }
  &__image {
    z-index:1;
    display:block;
    position:relative;
    transition: transform .3s ease-out;
    transform:translateX(0px);
    filter:drop-shadow(2px 0px 6px black);
    
    &.slide-open {
      transform:translateX(646px);
      transition: transform .3s ease-out;
    }

    &::before {
      content:'';
      position:absolute;
      width:50px;
      height:100px;
      left:-50px;
      top:50%;
      transform:translateY(-50%);
      background-color:$primary;
    }
  }
  &__heading {
    font-family:$heading-font;
    font-weight:300;
    font-size:1.5em;
    color:$primary-dark;
    justify-content:flex-end;
    display:flex;
    width:87%;
  }
  &__hidden-content {
    position:absolute;
    background-color:#e2e2e2;
    padding:20px 75px 20px 31px;
    top:0px;
    height:100%;
    box-shadow:inset -1px -14px 18px 0px rgba(0, 0, 0, .5);
    display:grid;
    grid-template-columns:1fr 1fr;
    grid-template-rows:92px 1fr;

    @include respond(mobileSmallPort) {
      padding:20px 40px 20px 31px;
    }

    .block-img {
      max-width:200px;
      img {
        object-fit:contain;
        mix-blend-mode:color-dodge;
      }
    }

    p {
      display:inline-flex;
      &:nth-of-type(3) {
        grid-column:2 span;
      }
    }
  }
  &__button {
    max-width:92px;
    max-height:63px;
    font-family:$heading-font;
    left:35%;
    top:-32px;
    position:absolute;
    transform:translateY(-50%);
    width:100%;
    height:100%;
    display:flex;
    align-items:center;
    justify-content:center;
    border:1px solid $color-black;
    cursor:pointer;
    font-size:1.5em;
    font-weight:300;
    @include respond(tabletLarge) {
       max-width:100px;
       max-height:100px;
       left:-100px;
       top:50%;
    }
  }
}
</style>