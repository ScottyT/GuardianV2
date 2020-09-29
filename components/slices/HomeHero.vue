<template>
	<div class="home-hero">
		<img :src="slice.primary.hero_image.url" :alt="slice.primary.hero_image.alt" />
		<div class="home-hero__wrapper">
			<!-- <h1 class="extra-large white--text mb-5 home-hero__wrapper--heading">{{ $prismic.asText(slice.primary.hero_cta_heading) }}</h1> -->
			<prismic-rich-text class="extra-large white--text home-hero__heading" wrapper="h1" :htmlSerializer="header" :field="slice.primary.hero_cta_heading" />
			<p class="white--text font-weight-light home-hero__wrapper--subheading">{{ $prismic.asText(slice.primary.hero_cta_subheading) }}</p>
		</div>
	</div>
</template>
<script>
import prismicDOM from 'prismic-dom'

const Elements = prismicDOM.RichText.Elements

const header =  function(type, element, children, content) {
	if (type === Elements.paragraph) {
		let result = ''
		const wrapperClassList = [element.label || '']
		result = `<span class="${wrapperClassList.join(' ')}">${content}</span>`;
		return result;
	}
	return null;
}
export default {
	name: "HomeHero",
	props: ["slice"],
	data() {
		return {
			header
		}
	}
};
</script>
<style lang="scss">
.home-hero {
	position: relative;
	@include respond(mobileSmall) {
		
	}
	@include respond(tabletLarge) {
		height:90vh;
	}
	@include respond(desktopSmall) {
		height:147vh;
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

	img {
		object-position:center center;
	}

	&::before {
		content: "";
		//background: linear-gradient(120deg, rgba(0, 0, 0, 0.88) 0%, rgba(0, 0, 0, 0.6) 45%, rgba(0, 0, 0, 0.4) 60%, rgba(0, 0, 0, 0) 100%);
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
	}

	&__wrapper {
		@include absCentered(43%, 50%);

		@include respond(mobileSmall) {
			max-width: 800px;
		}
		@include respond(mobileLarge) {
			max-width: 1000px;
			width: 100%;
		}

		&--heading {
			line-height: 1.3;
		}

		&--subheading {
			max-width: 100%;

			@include respond(mobileSmall) {
				max-width: 70%;
				font-size: 1.8rem;
			}
		}

		@media only screen and (max-width: 767px) {
			width: 95%;
			padding: 0 2rem;
		}
	}

	&__heading {
		line-height: 1.3;
	
		&::before {
			content:'';
			position:absolute;
			top:0;
			left:0;
			width:100%;
			height:67%;
			background:rgba($color-black, .6);
			z-index:-1;
		}
		&--primary {
			padding-left:25px;
		}
		&--secondary {
			display:block;
			background:rgba($color-black, .6);
			
			margin-left:15px;
			padding-left:50px;
			padding-right:15px;

			@include respond(mobileLarge) {
				width:95%;
			}

			@include respond(desktopSmall) {
				//max-width:898px;
				
			}
		}
	}
}
.hero {
}
</style>
