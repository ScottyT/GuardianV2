<template>
  <div class="content" v-if="loaded" id="home">
    <hero
      :hero="heroImage.url"
      :ctaHeading="heroCtaHeading[0].text | allCaps"
      :ctaSubheading="heroCtaSubheading[0].text"
      gradient="120deg, rgba(0,0,0,0.88) 0%, rgba(0,0,0,0.6) 45%, rgba(0,0,0,0.4) 60%,  rgba(0,0,0,0) 100%"
    />
    <div class="home-content">
      <side-navigation :links="this.sectionContent" :scroll="this.enableScroll"></side-navigation>
      <section
        v-for="(section, index) in sectionContent"
        :key="'section-' + index"
        :class="'section section--' + index"
        :id="'section--' + index"
        :ref="'section' + index"
      >
        <section-text :section="sectionContent[index]"></section-text>
      </section>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

const Hero = () => import(/*webpackPreload: true */ "@/components/Hero");
const SectionText = () =>
  import(/*webpackPreload: true */ "@/components/SectionText");
const SideNavigation = () =>
  import(/*webpackPreload: true */ "@/components/SideNavigation");

export default {
  name: "home",
  components: {
    Hero,
    SectionText,
    SideNavigation
  },
  data: () => {
    return {
      heroImage: [],
      heroCtaHeading: [
        {
          text: ""
        }
      ],
      heroCtaSubheading: [
        {
          text: ""
        }
      ],
      documentId: "",
      sectionContent: [],
      loaded: false,
      pageSections: [],
      mainNavLinks: [],
      selected: null,
      enableScroll: null,
      componentKey: 0
    };
  },
  mounted() {
    //console.log($(".page-sections__buttons").find("a"));
  },
  methods: {
    getContent() {
      this.$prismic.client.getSingle("home").then(document => {
        this.documentId = document.id;
        this.heroImage = document.data.hero_image;
        this.heroCtaHeading = document.data.hero_cta_heading;
        this.heroCtaSubheading = document.data.hero_cta_subheading;
        this.sectionContent = document.data.section;
        //this.pageSections = this.sectionContent;
        this.loaded = true;
        this.enableScroll = true;
      });
      // var sectionNodeList = document.querySelectorAll(".section");
      // var sections = Array.prototype.map.call(sectionNodeList, function(el) {
      //   return el;
      // });
      // this.pageSections = sections;
      // console.log(sections);
      // console.log(sectionNodeList);
      this.pageSections = this.$refs;
    }
  },
  created() {
    this.getContent();
  },
  beforeUpdate() {
    this.componentKey += 1;
  }
};
</script>
<style lang="scss">
.home-content {
  position: relative;
  padding: env(safe-area-inset-top) env(safe-area-inset-right)
    env(safe-area-inset-bottom) env(safe-area-inset-left);
  // @include respond(iPhoneX) {
  //   padding: env(safe-area-inset-top) env(safe-area-inset-right)
  //     env(safe-area-inset-bottom) env(safe-area-inset-left);
  // }
  @include respond(GalaxyS10Plus) {
    padding: env(safe-area-inset-top) env(safe-area-inset-right)
      env(safe-area-inset-bottom) env(safe-area-inset-left);
  }
}
</style>
