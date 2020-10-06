<template>
	<div class="page-hero">
		<img :src="slice.primary.hero.url" :alt="slice.primary.hero.alt" />
		<prismic-rich-text class="page-hero__button" :field="slice.primary.cta" />
	</div>
</template>
<script>
export default {
	name: "PageHero",
	props: ["slice"],
}
</script>
<style lang="scss">
.page-hero {
	position: relative;

	@include respond(mobileLargeMax) {
		height:399px;
	}

	img {
		@include respond(mobileLargeMax) {
			object-position:0px 0px;
		}
	}

	&__button {
		position:absolute;
		@include absCentered(93%, 24%);

		@include respond(mobileLarge) {
			@include absCentered(90%, 16%);
		}
	}
	// height: 300px;

	// @include respond(tabletLarge) {
	// 	height: 400px;
	// }

	// &::before {
	// 	content: "";
	// 	background: linear-gradient(120deg, rgba(0, 0, 0, 0.88) 0%, rgba(0, 0, 0, 0.6) 45%, rgba(0, 0, 0, 0.4) 60%, rgba(0, 0, 0, 0) 100%);
	// 	position: absolute;
	// 	width: 100%;
	// 	height: 100%;
	// 	top: 0;
	// }
}
</style>
