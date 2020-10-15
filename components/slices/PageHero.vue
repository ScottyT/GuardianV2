<template>
	<div :class="`page-hero page-hero--${$route.params.uid}`">
		<img :src="slice.primary.hero.url" :alt="slice.primary.hero.alt" />
		<prismic-rich-text :htmlSerializer="cta" :class="`page-hero__button page-hero__button--${slice.primary.button_position}`" :field="slice.primary.cta" />
	</div>
</template>
<script>
import prismicDOM from 'prismic-dom';
import linkResolver from '~/plugins/link-resolver.js';
const Elements = prismicDOM.RichText.Elements;

const cta = function(type, element, content, children) {
	if (type === Elements.hyperlink) {
		let result = '';
    const url = prismicDOM.Link.url(element.data, linkResolver);

    if (element.data.link_type === 'Document') {
      result = `<router-link to="${url}">${content}</router-link>`;
    } else {
			const target = element.data.target ? `target="'${element.data.target}'" rel="noopener"` : '';
      result = `<a class="${element.label}" href="${url}" ${target}>${content}</a>`;
    }
	}
}
export default {
	name: "PageHero",
	props: ["slice"],
	data() {
		return {
			cta
		}
	}
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

		&--right {
			@include absCentered(93%, 50%);
			@include respond(mobileSmallPort) {
				@include absCentered(93%, 60%);
			}
			
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

	&--about-us {
		img {
			object-position: 68% 0px;
			@include respond(mobileSmallPort) {
				
				object-position: -245px 0px;
			}
			@include respond(mobileLarge) {
				object-position:0px 0px;
			}			
		}
	}
}
</style>
