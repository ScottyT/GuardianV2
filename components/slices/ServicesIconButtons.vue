<template>
  <div class="services-buttons-wrapper">
    <div class="services-buttons" v-for="n in 5" :key="n">
      <div class="services-buttons__button">
        <h3>Roofing</h3>
        <span data-side="left" class="services-buttons__left-half"></span>
        <span data-side="right" class="services-buttons__right-half"></span>
      </div>
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
      let currentElm = null;
      button.onmouseover = function (event) {
        if (currentElm) return;
        console.log(event.target.dataset)
        if (event.target.dataset.side == "right") {
          button.querySelector('.services-buttons__left-half').classList.add('hovered')
        }
      }
      button.onmouseout = function (event) {
        button.querySelector('.services-buttons__left-half').classList.remove('hovered')
        if (!currentElm) return;
        let relatedTarget = event.relatedTarget;
        while (relatedTarget) {
          if (relatedTarget == currentElm) return;
          relatedTarget = relatedTarget.parentNode;
          
        }
        
        currentElm = null
      }
    }
  }
</script>
<style lang="scss">
  .services-buttons-wrapper {
    display: flex;
    height: 500px;
    width: 100%;
    justify-content: center;
    align-items: center;
    position: relative;
  }

  .services-buttons {
    height: 200px;
    width: 200px;
    //box-shadow:-12px 1px 7px -3px rgba(60,58,58,.45098);
    margin: 0 25px;
    position: relative;
    background-color: $color-white;
    box-shadow: none;

    h3 {
      text-align:center;
      width:100%;
    }

    &__button {
      perspective: 3000px;
      pointer-events: none;
      transform: rotateY(0);
      display: flex;
      height: 100%;
      transition: transform .3s ease-in;

      &:hover {
        transform: rotateY(30deg);
        transition: transform .3s ease-in;
      }

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
      box-shadow: none;
      transition: box-shadow .3s ease-in;
      &.hovered {
        box-shadow:23px 3px 12px -9px #0000008c;
      }

    }

    &__left-half {
      box-shadow: none;
      transition: box-shadow .3s ease-in;

      &:hover {
        transition: box-shadow .3s ease-in;
        box-shadow: -12px 1px 7px -3px rgba(60, 58, 58, .45098);
      }

      &.hovered {
        transition: box-shadow .3s ease-in;
        box-shadow: -23px 0px 12px -9px #0000008c;
      }
    }

    &::after {
      position: absolute;
      content: "";
      z-index: -1;
      bottom: 15px;
      right: 8px;
      top: 80%;
      width: 50%;
      box-shadow: -5px 20px 16px 0px #0b0c0c;
      transform: rotate(10deg);
    }

    &::before {
      content: '';
      position: absolute;
      bottom: 15px;
      z-index: -1;
      left: 8px;
      top: 80%;
      width: 50%;
      background: #111;
      box-shadow: -5px 20px 16px 0px #0b0c0c;
      transform: rotate(-10deg);
    }
  }
</style>