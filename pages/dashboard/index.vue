<template>
  <div class="user-area">
 
    <h1>User Area</h1>

  </div>
</template>
<script>
export default {
  layout: 'dashboard',
}
</script>
<style lang="scss">
.slide-right {
  opacity:0;
  transform:translateX(200px);
  transition: all .2s ease-in;
}
.slide-left {
  opacity:1;
  transform:translateX(-200px)
}
</style>