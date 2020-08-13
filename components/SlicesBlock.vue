<template>
	<div class="slices-wrapper">
		<section v-for="(slice, index) in slices" :key="index" class="section">
			<template v-if="slice.slice_type === 'homehero'">
				<home-hero :slice="slice" />
			</template>
			<template v-else-if="slice.slice_type === 'pagehero'">
				<page-hero :slice="slice" />
			</template>
			<template v-else-if="slice.slice_type === 'twocolcardlist'">
				<two-col-cards :slice="slice" />
			</template>
			<template v-else-if="slice.slice_type === 'offsetsectiontext'">
				<offset-section-text :slice="slice" />
			</template>
		</section>
	</div>
</template>
<script>
export default {
	name: "slices-block",
	props: {
		slices: Array,
	},
	components: {
		TwoColCards: () => import("./slices/TwoColCardList.vue"),
	},
}
</script>
<style lang="scss">
html {
	@include respond(mobileSmall) {
		font-size: 56.5% !important;
	}

	@include respond(mobileLarge) {
		font-size: 58.5% !important;
	}

	@include respond(tabletLarge) {
		font-size: 62.5% !important;
	}
}
body {
	font-size: 18px;
}
h1 {
	font-family: $heading-font;
	letter-spacing: 1px;
	@include respond(desktopSmall) {
		font-size: 4em;
	}

	&.extra-large {
		@include respond(tabletLarge) {
			font-size: 6em;
		}
	}
}
h2 {
	font-size: 1.7em;
	line-height: 1.3;
	font-family: $heading-font;
	@include respond(mobileLarge) {
		font-size: 2em;
	}
	@include respond(tabletLarge) {
		font-size: 2.5em;
	}
	@include respond(desktopSmall) {
		font-size: 3em;
	}
}
ul {
	list-style: none;
}
img {
	width: 100%;
	height: 100%;
	object-fit: cover;
}
.section {
	position: relative;
	margin: 0 0 4rem 0;
	height: auto;
}
</style>
