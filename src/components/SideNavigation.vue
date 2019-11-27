<template>
  <nav class="page-sections">
    <ul class="page-sections__buttons" id="dots-wrapper">
      <a
        class="page-sections__buttons--dots"
        v-for="(section, index) in links"
        :key="index"
        :data-link="'#section--' + index"
        @click="buttonSelect($event)"
      ></a>
      <div class="page-sections__marker">
        <div class="sr-only">Section marker</div>
      </div>
    </ul>
  </nav>
</template>
<script>
import $debounce from "lodash.debounce";
import JQuery from "jquery";
var $ = JQuery;
export default {
  name: "SideNavigation",
  props: {
    links: Array,
    scroll: Boolean
  },
  mounted() {
    // var pageSections = document.getElementById("dots-wrapper");

    const sections = document.querySelectorAll("section");
    var intersectionOptions = {
      root: null, // use the viewport
      rootMargin: "0px",
      threshold: 0.55
    };
    function handleIntersect(entries) {
      entries.forEach(entry => {
        var pageSection = entry.target.attributes[0].nodeValue;
        if (entry.intersectionRatio > 0.55) {
          var id = entry.target.attributes[0].nodeValue;
          document
            .querySelector(`[data-link='#${id}']`)
            .classList.add("current");
        } else {
          document
            .querySelector(`[data-link='#${pageSection}']`)
            .classList.remove("current");
        }
      });
    }
    function createObserver() {
      var observer = new IntersectionObserver(
        handleIntersect,
        intersectionOptions
      );
      sections.forEach(section => {
        observer.observe(section);
      });
    }
    createObserver();

    document.addEventListener("scroll", this.scrollingMarker);
  },
  methods: {
    buttonSelect: function(event) {
      var sectionElm = document.querySelectorAll(
        event.target.attributes[0].nodeValue
      )[0];
      var topSection = sectionElm.offsetTop + window.innerHeight / 2;
      window.scrollTo({
        top: topSection,
        left: 0,
        behavior: "smooth"
      });
    },
    scrollingMarker: $debounce(function() {
      const sections = document.querySelectorAll("section");
      sections.forEach(item => {
        var $sidebar = $(".page-sections__marker"),
          $window = $(window);
        var section = $("#" + item.attributes[0].nodeValue)[0];
        var sectionOffset = section.getBoundingClientRect();
        if (
          $window.scrollTop() <= section.offsetHeight / 2 + section.offsetTop &&
          $window.scrollTop() > 496
        ) {
          $sidebar.stop().animate({
            top: $window.scrollTop() / 15 - 15
          });
        } else if (
          window.scrollY >=
          section.offsetTop + sectionOffset.height / 2
        ) {
          $sidebar.stop().animate({
            top: 125
          });
        } else {
          $sidebar.stop().animate({
            top: 5
          });
        }
        //movingMarker("#" + link.attributes[0].nodeValue);
      });
    }, 100)
  },
  destroyed() {
    document.removeEventListener("scroll", this.scrollingMarker);
  }
};
</script>
<style lang="scss">
#dots-wrapper {
  @media only screen and (min-width: $lg) {
    padding-left: 0;
  }
}
.page-sections {
  position: sticky;
  position: -webkit-sticky;
  z-index: 10;
  top: 55px;
  display: none;

  @include respond(tabletLarge) {
    display: block;
  }

  @include respond(iPhoneX) {
    display: none;
  }

  &__buttons {
    position: absolute;
    height: 130px;
    top: 155px;
    right: 35px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;

    @media only screen and (max-width: 1200px) {
      right: 27px;
    }

    @include respond(iPhoneX) {
      right: 15px;
    }

    @include respond(mobileLand) {
      top: 122px;
    }

    &--dots {
      background-color: $grey;
      width: 20px;
      height: 10px;
      z-index: 4;

      &:hover {
        background-color: $primary-dark;
      }

      &.current {
        background-color: $tertiary-light;
      }
    }
  }

  &__marker {
    width: 70px;
    height: 1px;
    border-top: 1px solid $secondary-dark;
    position: absolute;
    top: 5px;
    left: -30px;

    @media screen and (max-width: 1200px) {
      width: 40px;
      left: 12px;
      border-top-width: 2px;
    }

    @include respond(mobileLand) {
      display: none;
    }
  }
  button {
    z-index: 10;
  }
}
</style>
