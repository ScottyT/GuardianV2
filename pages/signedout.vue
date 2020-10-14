<template>
  <section class="section section--signedout">
    <h1>You have signed out</h1>
    <h2>You will be redirected in 3 seconds</h2>
  </section>
</template>
<script>
export default {
  mounted() {
    this.$nextTick(() => {
      var delay = 3000

      setTimeout(() => {
        this.$router.go(-1)
      }, delay)
    })
  }
}
</script>
<style lang="scss">
.section {
  &--signedout {
    padding-top:100px;
  }
}
</style>