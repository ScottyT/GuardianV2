<template>
  <div :class="slice.primary.class_wrapper">
    <a @click="isActive = !isActive" v-if="slice.primary.button != null" role="button" class="button button--red button--expand">More</a>
    <lazy-image :source="slice.primary.image.url" :imageClass="`${slice.primary.image_class} ${isActive ? 'opened' : ''}`" :alt="slice.primary.image.alt" />
  </div>
</template>
<script>
export default {
  name: 'SingleImage',
  props:['slice'],
  data() {
    return {
      isActive: false
    }
  }
}
</script>
<style lang="scss">
.sign-wrapper {
  background-image:url('https://images.prismic.io/guardianrestoration/d7392d72-80b8-4d1c-bc14-6b693ef5d995_Brick+Wall+1920x950.png?auto=compress,format');

  &__sign {
    height:auto;
    max-width:1200px;
    margin:auto;
    transform:translateY(-100%);
    transition:transform .4s ease-in;
    &.opened {
      transform:translateY(-11%);
      transition:transform .8s cubic-bezier(0.7, -0.01, 0.33, 1.45);
    }
    img {
      object-fit:contain;
    }
  }
}
</style>