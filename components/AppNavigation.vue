<template>
	<v-app-bar :height="navbarHeight" hide-on-scroll elevate-on-scroll class="navigation" :dark="!scrolledDown" fixed>
		<nuxt-link class="navigation__site-title" to="/">Gaurdian Restoration</nuxt-link>
		<v-toolbar-items class="navigation">
			<ul class="v-toolbar__items">
				<li v-for="(menuLink, index) in $store.state.header.menu_links" :key="index" class="navigation__menu-item">
					<prismic-link :field="menuLink.link">{{ $prismic.asText(menuLink.label) }}</prismic-link>
				</li>
			</ul>
		</v-toolbar-items>
	</v-app-bar>
</template>
<script>
import $debounce from "lodash.debounce"
export default {
	name: "AppNavigation",
	components: {},
	data() {
		return {
			scrolledDown: false,
			menuContent: [],
			menuLinks: [],
			hasClass: false,
		}
	},
	mounted() {
		window.addEventListener("scroll", this.isScrolling)
	},
	computed: {
		navbarHeight() {
			var viewportWidth = this.$vuetify.breakpoint.width

			switch (true) {
				case viewportWidth <= 500:
					return "65px"

				case viewportWidth <= 768:
					return "58px"

				case viewportWidth < 960:
					return "56px"

				default:
					return "65px"
			}
		},
	},
	methods: {
		isScrolling: $debounce(function() {
			this.scrolledDown = window.scrollY > 50 || this.hasClass
		}, 100),
	},
	created() {
		//this.isScrolling()
	},
}
</script>
<style lang="scss">
.v-app-bar {
	&--is-scrolled {
		background-color: black;
	}
}
.v-toolbar {
	&__content {
		padding: 4px 40px !important;
		justify-content: space-between;

		@include respond(mobileSmall) {
			flex-direction: column;
		}

		@include respond(mobileSmallLand) {
			flex-direction: row;
		}

		@include respond(mobileLarge) {
			flex-direction: row;
		}
	}

	&__items {
		flex: 0 0 40%;
		max-width: 1170px;
		justify-content: space-between;

		@include respond(mobileSmall) {
			height: 32px;
		}
	}
}
.navigation {
	@include respond(mobilePort) {
		height: 79px;
	}
	&__site-title {
		color: inherit;

		@include respond(mobilePort) {
			padding: 7px 0;
		}
	}

	&__menu-item {
		min-width: 135px;
		height: 100%;
		display: flex;
		background-image: linear-gradient(120deg, transparent 0%, transparent 50%, rgba($color-white, 0.2) 50%);
		background-size: 245%;
		transition: all 0.4s ease-in;
		justify-content: center;

		&:hover {
			background-position: 100%;
		}

		a {
			padding: 0 16px;
			display: flex;
			align-items: center;
			font-weight: 400;
			letter-spacing: 0.15rem;
			text-transform: uppercase;
			position: relative;
			font-size: 14px;
			color: inherit !important;

			@include respond(mobileSmallLand) {
				font-size: 14px;
			}

			@include respond(mobileLarge) {
				font-size: 14px;
			}

			&::before {
				content: "";
				position: absolute;
				top: 0;
				left: 0;
				bottom: 0;
				right: 0;
				opacity: 0;
				transition: opacity 0.3s;
			}
		}
	}

	.router-link-exact-active {
		&::before {
			content: "";
			background-color: $color-white;
			opacity: 0.12;
			height: 100%;
			width: 100%;
		}
	}
}
</style>
