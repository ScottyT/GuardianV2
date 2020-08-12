<template>
	<div class="offset-section-text">
		<div class="offset-section-text__row" v-for="(section, i) in slice.items" :key="i">
			<div :class="`offset-section-text__content offset-section-text__content--${(i + 1) % 2 === 0 ? 'left' : 'right'}`">
				<prismic-rich-text :htmlSerializer="heading" :field="section.sectionheading" v-if="section.sectionheading.length > 0" />
				<p>{{ $prismic.asText(section.sectiontext) }}</p>
			</div>
			<div :class="`offset-section-text__image offset-section-text__image--${(i + 1) % 2 === 0 ? 'left' : 'right'}`">
				<img :src="section.sectionimage.url" :alt="section.sectionimage.alt" />
			</div>
		</div>
	</div>
</template>
<script>
import prismicDOM from "prismic-dom";
const Elements = prismicDOM.RichText.Elements;

const heading = function (type, element, content, children) {
	switch (type) {
		case Elements.paragraph:
			return `<p class="offset-section-text__content--subheading">${children.join("")}</p>`;
		default:
			return null;
	}
};
export default {
	name: "OffsetSectionText",
	props: ["slice"],
	data() {
		return {
			heading,
		};
	},
};
</script>
<style lang="scss">
.offset-section-text {
	&__row {
		display: flex;
		flex-wrap: wrap;

		@include respond(tabletLarge) {
			padding: 0 0 3rem 0;
		}

		&:first-child {
			.offset-section-text__content {
				margin-top: -10rem;
				background-color: $color-white;
				position: relative;
			}
		}
		&:nth-of-type(even) {
			flex-direction: row-reverse;
		}
	}
	&__content {
		width: 50%;
		display: grid;
		place-content: center;

		@include respond(desktopSmall) {
			grid-template-columns: 500px;
		}
		&--left {
			padding: 32px;
		}
		&--right {
			padding: 0 28px 0 40px;
		}
		&--subheading {
			font-size: 1.2em;
			text-transform: uppercase;
		}
	}
	&__image {
		width: 50%;
		position: relative;

		&--right {
			max-width: 625px;
			min-height: 250px;
			left: -20px;
			margin-top: 7rem;
		}
		&--left {
			max-width: 555px;
			right: -10px;
			max-height: 333px;
		}
	}
}
</style>
