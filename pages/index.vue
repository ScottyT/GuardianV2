<template>
	<slices-block :slices="slices" />
	<!-- <div v-if="loaded" id="home" class="content">
		
		<div class="home-content">
			<side-navigation
        :links="this.sectionContent"
        :scroll="this.enableScroll"
      ></side-navigation>
			<section v-for="(section, index) in sectionContent" :id="'section--' + index" :key="'section-' + index" :ref="'section' + index" :class="'section section--' + index">
				<section-text :section="sectionContent[index]"></section-text>
			</section>
		</div>
	</div> -->
</template>

<script>
export default {
	name: "Index",
	components: {},
	layout: "homepage",
	data: () => {
		return {
			componentKey: 0,
		};
	},
	created() {
		this.getContent();
	},
	beforeUpdate() {
		this.componentKey += 1;
	},
	async asyncData({ $prismic, error }) {
		const homepage = (await $prismic.api.getSingle("home")).data;
		return {
			slices: homepage.body,
		};
	},
	methods: {
		getContent() {
			// this.$prismic.client.getSingle('home').then((document) => {
			//   this.documentId = document.id
			//   this.heroImage = document.data.hero_image
			//   this.heroCtaHeading = document.data.hero_cta_heading
			//   this.heroCtaSubheading = document.data.hero_cta_subheading
			//   this.sectionContent = document.data.section
			//   this.loaded = true
			//   this.enableScroll = true
			// })
			// this.pageSections = this.$refs
		},
	},
};
</script>
<style lang="scss">
.home-content {
	position: relative;
	padding: env(safe-area-inset-top) env(safe-area-inset-right) env(safe-area-inset-bottom) env(safe-area-inset-left);
	@include respond(iPhoneX) {
		padding: env(safe-area-inset-top) env(safe-area-inset-right) env(safe-area-inset-bottom) env(safe-area-inset-left);
	}
	@include respond(GalaxyS10Plus) {
		padding: env(safe-area-inset-top) env(safe-area-inset-right) env(safe-area-inset-bottom) env(safe-area-inset-left);
	}
}
</style>
