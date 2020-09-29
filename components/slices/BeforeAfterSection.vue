<template>
  <div class="offset-section-text">
    <div class="offset-section-text__row">
      <div class="offset-section-text__image before-after-block">
        <div class="before-after-block__image before-after-block__image--before">
          <img :src="slice.primary.before_image.url" :alt="slice.primary.before_image.alt" />
        </div>
        <div class="before-after-block__image before-after-block__image--after">
          <img :src="slice.primary.after_image.url" :alt="slice.primary.after_image.alt" />
        </div>
      </div>
      <div class="offset-section-text__content">
        <p>{{$prismic.asText(slice.primary.text)}}</p>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'BeforeAfterSection',
  props: ['slice']
}
</script>
<style lang="scss">
.before-after-block {
  position:relative;
  overflow:hidden;
  height:400px;

  &__image {
    position:absolute;
    max-width:555px;
    overflow:hidden;
    right:-10px;
  }
}
</style>