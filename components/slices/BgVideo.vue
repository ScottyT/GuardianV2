<template>
  <video autoplay muted loop class="page-hero__video" ref="bgVideo">
    <source src="~/assets/video/Contact Us on Stage CROP w sharpen-Orig.mp4" type="video/mp4" />
  </video>
</template>
<script>
export default {
  name: "BgVideo",
  props: ['slice'],
  mounted() {
    const config = {
      root: null,
      rootMargin: '0px',
      threshold: 0
    }
    const videoBg = document.querySelectorAll('.page-hero__video');
    var observer = new IntersectionObserver((entries, self) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          entry.target.play()
        } else {
          entry.target.pause()
        }
      })
    }, config);
    videoBg.forEach((video) => {
      observer.observe(video)
    })
  }
}
</script>
<style lang="scss">
.page-hero {
  &__video {
    display:block;
  }
}
</style>