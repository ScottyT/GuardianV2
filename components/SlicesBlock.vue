<template>
	<div class="slices-wrapper">
		<section class="home-hero-wrapper" v-for="(slice, i) in homeHeroSlice" :key="i">
			<template v-if="slice.slice_type === 'homehero'">
				<home-hero :slice="slice" />
			</template>
		</section>
		<div class="slices-wrapper__body-wrapper">
		<section v-for="(slice, index) in slices" :key="index" class="section">			
			<template v-if="slice.slice_type === 'pagehero'">
				<page-hero :slice="slice" />
			</template>
			<template v-else-if="slice.slice_type === 'twocolcardlist'">
				<two-col-cards :slice="slice" />
			</template>
			<template v-else-if="slice.slice_type === 'offsetsectiontext'">
				<offset-section-text :slice="slice" />
			</template>
			<template v-else-if="slice.slice_type === 'contactus'">
				<contact-us :slice="slice" />
			</template>
			<template v-else-if="slice.slice_type === 'rollupsection'">
				<roll-up-section :slice="slice" />
			</template>
			<template v-else-if="slice.slice_type === 'beforeandafter'">
				<before-after-section :slice="slice" />
			</template>
		</section>
		</div>
	</div>
</template>
<script>
export default {
	name: "slices-block",
	props: {
		slices: Array,
	},
	computed: {
		homeHeroSlice() {
			var slicesArr = this.slices
			const heroSliceIndex = slicesArr.filter((slice) => {
				return slice.slice_type == "homehero"
			})
			return heroSliceIndex
		},
		// filteredSlices() { 
		// 	var removeElement = this.slices.shift()
		// 	return this.slices
		// }
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
		font-size:2.7em;
		@include respond(mobileLarge) {
			font-size:3.5em;
		}
		
		@include respond(tabletLarge) {
			font-size:5em;
		}
		@include respond(desktopSmall) {
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
.slices-wrapper {
	position:relative;

	&__body-wrapper {
		background-image:url('https://images.prismic.io/guardianrestoration/9d94a63c-fe7f-461b-bc1f-c5b446dccd79_Triangles+BG+for+Guardian+home.png?auto=compress,format');
		background-repeat:repeat;
		padding-top:30px;
		background-position-y:-305px;
	}
}
.section {
	position: relative;
	height: auto;
}
.form {
	&__input {
		padding: 6px 12px;
		border: 1px solid $color-black;
		&:not(:last-child) {
			margin-bottom: 5px;
		}
	}
}
.block-heading {
	position:relative;
}
</style>
