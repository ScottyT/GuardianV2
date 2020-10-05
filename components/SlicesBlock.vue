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
				<template v-else-if="slice.slice_type === 'fullwidthimage'">
					<full-width-image :slice="slice" />
				</template>
			</section>
		</div>
		<div class="slices-wrapper__service-buttons">
			<section v-for="slice in serviceButtonSlice" :key="slice.id">
				<template v-if="slice.slice_type === 'servicesiconbuttons'">
					<services-icon-buttons :slice="slice" />
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
		serviceButtonSlice() {
			var slicesArr = this.slices
			return slicesArr.filter((slice) => {
				return slice.slice_type == "servicesiconbuttons"
			})
		}
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
.v-application--wrap {
	overflow:hidden;
}
.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  border: 0;
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
		font-size:2.5em;
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
a {
	&.button {
		display:inline-block;
		padding:4px 14px;
		max-width:100px;
		width:100%;
		color:$color-white;

		&--red {
			background: rgb(251,132,132);
			background: linear-gradient(90deg, rgba(251,132,132,1) 0%, rgba(211,83,83,1) 50%);
			transition:box-shadow .3s ease-in;

			&:hover {
				box-shadow:inset 0 0 0 4px $color-black;
				transition:box-shadow .3s ease-in;
			}
		}
		&--services {
			color:$color-white;
			position:absolute;
			left:50%;
			transform:translateX(-50%) translateY(-5px);
			font-size:1.2em;
			text-transform:uppercase;
			font-family:$heading-font;
			z-index:-1;
			bottom:23px;
			opacity:0;
			transition:opacity .5s ease-in, transform .5s ease-in;
		}
	}
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
		background-size:contain;

		@include respond(mobileLarge) {
			background-position-y:-305px;
			background-size:cover;
		}
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
