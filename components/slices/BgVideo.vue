<template>
  
  <div class="page-hero__video">
    <span class="overlay">
      <img :class="buffer || $vuetify.breakpoint.width < 959 ? '' : 'hidden'" src="https://images.prismic.io/guardianrestoration/d04a9219-a032-420e-b90f-0b7bed7124ac_Contact+Us+STILL.png?auto=compress,format" />
    </span>
    <video v-show="$vuetify.breakpoint.width > 960" autoplay muted loop ref="bgVideo" class="page-hero__video-wrapper">
      <source src="~/assets/video/Contact Us on Stage CROP w sharpen-Orig.mp4" type="video/mp4" />
    </video>
    <!-- <client-only>
      <youtube @ready="ready" @playing="playing" @buffering="buffering" :video-id="slice.primary.videoid" player-width="100%" player-height="600px" :player-vars="{autoplay:1, rel:0, loop:1, mute:1, playlist:slice.primary.videoid, origin: 'http://localhost:3000'}"></youtube>
    </client-only> -->
    
    <!-- <iframe v-show="$vuetify.breakpoint.width > 960" width="100%" height="100%" src="https://www.youtube.com/embed/YhpFg9Fippk?enablejsapi=1&rel=0&autoplay=1&loop=1&playlist=YhpFg9Fippk&mute=1" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope;" allowfullscreen></iframe> -->
  </div>
</template>
<script>
import { getIdFromURL, getTimeFromURL } from 'vue-youtube-embed'
export default {
  name: "BgVideo",
  props: ['slice'],
  data() {
    return {
      videoId: 'videoId',
      buffer: false
    }
  },
  methods: {
    ready(event) {
      this.player = event.target
      
      console.log(this.player)
    },
    playing(event) {
      console.log("video is playing:", event)
      this.buffer = false
    },
    buffering(event) {
      console.log("buffering")
      this.buffer = true
    },
    play() {
      this.player.playVideo()
    }
  },
  mounted() {
    // const config = {
    //   root: null,
    //   rootMargin: '0px',
    //   threshold: 0
    // }
    // const videoBg = document.querySelectorAll('.page-hero__video');
    // var observer = new IntersectionObserver((entries, self) => {
    //   entries.forEach((entry) => {
    //     if (entry.isIntersecting) {
    //       entry.target.play()
    //     } else {
    //       entry.target.pause()
    //     }
    //   })
    // }, config);
    // videoBg.forEach((video) => {
    //   observer.observe(video)
    // })
  }
}
</script>
<style lang="scss">
.page-hero {
  &__video-wrapper {
    position:absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    display: block;
    height: auto;
    width: 100%;
  }
  &__video {
    display:block;
    height:190px;
    position:relative;
    @include respond(mobileSmallPort) {
      height:350px;
    }

    @include respond(tabletLarge) {
      height:30vw;
    }
    
    .overlay {
      position:absolute;
      width:100%;
      height:100%;
      //z-index:2;
      top:32px;
      @include respond(mobileSmallPort) {
        top:0;
      }

      img {
        &.hidden {
          display:none;
        }
      }
    }
    
    iframe {
      height: calc(100% + 250px);
      width: 100%;
      position: absolute;
      top: 0;
      left: 0;
      bottom: 0;
      pointer-events: none;
      margin-top: -89px;
    }
    // @media (min-aspect-ratio: 16/9) {
    //   width:100%;
    //   height:auto;
    // }
    // @media (max-aspect-ratio: 16/9) {
    //   width:auto;
    //   height:100%;
    // }
  }
}
</style>