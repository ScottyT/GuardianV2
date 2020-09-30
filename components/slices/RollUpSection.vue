<template>
  <div class="services-featured" :style="`background-image:url(${slice.primary.bg_image.url}`">
    <div :class="`services-featured__col services-featured__col--${i}`" :data-col="i" v-for="(column, i) in slice.items" :key="i">
      <div class="services-featured__heading">{{$prismic.asText(column.col_title)}}</div>
      <div class="services-featured__subheading">{{$prismic.asText(column.col_subtitle)}}</div>
      <p class="services-featured__text">{{$prismic.asText(column.col_text)}}</p>
    </div>
  </div>
</template>
<script>
import anime from 'animejs';
import $debounce from 'lodash.debounce';
export default {
  name: "RollUpSection",
  props: ['slice'],
  mounted() {
    function hoverAnim() {
      var classTarget = this.classList[1]
      anime({
        targets: `.${classTarget}`,
        translateY: 0,
        easeing: 'easeInOut'
      })
      console.log(classTarget)
    }

    var serviceCols = document.querySelectorAll('.services-featured__col');
    serviceCols.forEach((el) => {
      console.log(el)
      el.addEventListener("mouseover", hoverAnim, false)
    })
  }
}
</script>
<style lang="scss">
.services-featured {
  height:700px;
  width:100%;
  max-width:1600px;
  margin:0 auto 50px;
  background-size:cover;
  color:$color-white;
  // display:grid;
  // grid-template-columns: repeat(4, 1fr);
  overflow:hidden;
  display:flex;
  justify-content: space-around;

  &__heading {
    font-family:$heading-font;
    line-height:1.1;
    z-index:1;
    font-size:2em;
    text-transform: uppercase;
    @include respond(desktopSmall) {
      font-size:2.7em;
    }
  }
  &__subheading {
    font-size:1.5em;
    font-family:$heading-font;
    z-index:1;
    
  }
  &__text {
    margin-top:10px;
    z-index:1;
    font-size:.9em;
    @include respond(desktopSmall) {
      font-size:1em;
    }
  }
  &__col {
    padding:10px;   
    display:flex;
    flex-direction:column;
    align-items:center;
    width:24%;
    position:relative;
    border-left:11px solid $primary-dark;
    transform:translateY(83%);
    transition:all .5s cubic-bezier(0.65, 0.06, 0.29, 0.93);

    &::before {
      z-index:0;
      content:'';
      position: absolute;
      width:100%;
      height:100%;
      top:0;
      left:0;
      opacity:0;
      background-color:rgba($color-grey, 1);
      transition:all .5s cubic-bezier(0.65, 0.06, 0.29, 0.93);
    }

    // &:hover {
    //   transition:all .5s cubic-bezier(0.65, 0.06, 0.29, 0.93);
    //   transform:translateY(0);
    //   &::before {
    //     opacity:.9;
    //   }
    // }
  }
}
</style>