<template>
  <div class="services-buttons-wrapper">
    <div class="services-buttons" v-for="n in 6" :key="n">
      <div class="services-buttons__button">
        <div class="services-buttons__button--icon">
          <img class="icon icon--centered"
            src="https://images.prismic.io/guardianrestoration/d1fe609f-2a2a-4c68-900d-0788b310fc3f_Roofing+Centered.png?auto=compress,format" />
          <img class="icon icon--left-pushed hidden"
            src="https://images.prismic.io/guardianrestoration/375399f0-e528-4740-85bb-4f27e0b04e77_Roofing+L+Pushed+down.png?auto=compress,format" />
          <img class="icon icon--right-pushed hidden"
            src="https://images.prismic.io/guardianrestoration/7741a1d9-58b7-4500-8e22-bf978381bd60_Roofing+R+Pushed+down.png?auto=compress,format" />
            <a class="button button--services button--red">Explore</a>
        </div>
        <span data-side="left" class="services-buttons__half services-buttons__left-half"></span>
        <span data-side="right" class="services-buttons__half services-buttons__right-half"></span>
        
      </div>
    </div>
    <div class="gr-logo-col">
      <img src="https://images.prismic.io/guardianrestoration/b0aa757b-f95c-4a67-a29d-197e68e6a0b1_G+big+Guardian.png?auto=compress,format" />
      <a class="button button--red">More</a>
    </div>
  </div>
</template>
<script>
  import anime from 'animejs'
  export default {
    name: "ServicesIconsButtons",
    props: ['slice'],
    mounted() {
      var button = document.querySelector('.services-buttons-wrapper');
      var icon = document.querySelector('.services-buttons__button--icon');
      let currentElm = null;
      button.onmouseover = function (event) {
        if (currentElm) return;
        let target = event.target.closest('div.services-buttons__button')
        //let half = target.querySelector('.services-buttons__half')
        //console.log(half)
        if (!target) return;
        if (!button.contains(target)) return;
        currentElm = target
        console.log(event.clientX)
        if (event.target.dataset.side == "right") {          
          target.querySelector('.icon--left-pushed').classList.add('hidden')
          target.querySelector('.icon--right-pushed').classList.remove('hidden')
          target.querySelector('.icon--centered').classList.add("hidden")
        }
        if (event.target.dataset.side == "left") {
          target.querySelector('.icon--left-pushed').classList.remove('hidden')
          target.querySelector('.icon--right-pushed').classList.add('hidden')
          target.querySelector('.icon--centered').classList.add("hidden")
        }
      }
      button.onmouseout = function (event) {
        if (!currentElm) return;
        let relatedTarget = event.relatedTarget;
        while (relatedTarget) {
          if (relatedTarget == currentElm) return;
          relatedTarget = relatedTarget.parentNode;
        }
        if (event.target.dataset.side == "right") {
          console.log("mouse out")
          currentElm.querySelector('.icon--left-pushed').classList.add('hidden')
          currentElm.querySelector('.icon--right-pushed').classList.add('hidden')
          currentElm.querySelector('.icon--centered').classList.remove("hidden")
        }
        if (event.target.dataset.side == "left") {        
          currentElm.querySelector('.icon--left-pushed').classList.add('hidden')
          currentElm.querySelector('.icon--right-pushed').classList.add('hidden')
          currentElm.querySelector('.icon--centered').classList.remove("hidden")
        }
        currentElm = null
      }
    }
  }
</script>
<style lang="scss">
  .services-buttons-wrapper {
    display: grid;
    column-gap:40px;
    row-gap:40px;
    padding:20px 0 70px;
    width: 100%;
    justify-content: center;
    align-items: center;
    position: relative;
    grid-template-columns:200px;
    @include respond(mobileSmallPort) {
      grid-template-columns:repeat(2, 200px);
      grid-template-rows:repeat(4, 230px);
      row-gap:10px;
    }
    @include respond(mobileLarge) {
      grid-template-columns:repeat(2, 222px) 200px;
      grid-template-rows:repeat(3, 231px);
    }
    
    @include respond(desktopSmall) {
      grid-template-columns:repeat(3, 200px) 250px;
      grid-template-rows:repeat(2, 260px);
    }

    .gr-logo-col {      
      text-align:center;

      @include respond(mobileLarge) {
        grid-column:3;
        grid-row:1/2 span;
      }

      @include respond(desktopSmall) {
        grid-column:4;
      }
    }
  }

  .services-buttons {
    //box-shadow:-12px 1px 7px -3px rgba(60,58,58,.45098);
    position: relative;
    //background-color: $color-white;
    box-shadow: none;

    h3 {
      text-align: center;
      width: 100%;
    }

    &__button {
      //perspective: 3000px;
      //transform: rotateY(0);
      display: flex;
      height: 110%;
      transition: all .3s ease-in;

      &:hover {
        .button--services {
          opacity:1;
          transform:translateX(-50%) translateY(62px);
          transition:all .5s ease-in;
        }
      }

      &--icon {
        height: 200px;
        width: 200px;
        z-index:0;
      }

      .icon {
        position: absolute;
        top: 0;
        width: 100%;
        height: 100%;
        opacity:1;
        transition:all .3s ease-in;
        
      }
      .hidden {
        opacity:0;
        transition:all .3s ease-in;
      }

      // &:hover {
      //   transform: rotateY(30deg);
      //   transition: transform .3s ease-in;
      // }

      .services-buttons__right-half {
        pointer-events: auto;
      }

      .services-buttons__left-half {
        pointer-events: auto;
      }
    }

    &__left-half,
    &__right-half {
      width: 50%;
      height: 100%;
      display: inline-block;
      position: absolute;
      top: 0;
    }

    &__right-half {
      right: 0;
      float: right;
    }
  }
</style>