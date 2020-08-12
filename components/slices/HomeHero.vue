<template>
	<div class="home-hero">
		<img :src="slice.primary.hero_image.url" :alt="slice.primary.hero_image.alt" />
		<div class="home-hero__wrapper">
			<h1 class="extra-large white--text mb-5 home-hero__wrapper--heading">{{ $prismic.asText(slice.primary.hero_cta_heading) }}</h1>
			<p class="white--text font-weight-light home-hero__wrapper--subheading">{{ $prismic.asText(slice.primary.hero_cta_subheading) }}</p>
		</div>
	</div>
</template>
<script>
export default {
	name: "HomeHero",
	props: ["slice"],
};
</script>
<style lang="scss">
.home-hero {
	position: relative;
	@include respond(mobileSmall) {
		height: 800px;
	}

	@include respond(mobileSmallLand) {
		height: calc(100vh - 100px);
	}

	@include respond(iPhoneX) {
		height: calc(100vh - 100px);
	}
	@include respond(GalaxyS10Plus) {
		height: calc(100vh - 100px);
	}

	&::before {
		content: "";
		background: linear-gradient(120deg, rgba(0, 0, 0, 0.88) 0%, rgba(0, 0, 0, 0.6) 45%, rgba(0, 0, 0, 0.4) 60%, rgba(0, 0, 0, 0) 100%);
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
	}

	&__wrapper {
		max-width: 800px;
		width: 79%;
		position: absolute;

		@include respond(mobileSmall) {
			top: 55%;
			transform: translateY(-50%);
		}
		@include respond(mobileLarge) {
			max-width: auto;
			@include absCentered(50%, 50%);
			width: 89%;
		}

		&--heading {
			line-height: 1.3;
		}

		&--subheading {
			max-width: 100%;

			@include respond(mobileLarge) {
				max-width: 70%;
				font-size: 1.8rem;
			}
		}

		@media only screen and (max-width: 767px) {
			max-width: 100%;
			width: auto;
			padding: 0 5rem;
		}
	}
}
.hero {
}
</style>
