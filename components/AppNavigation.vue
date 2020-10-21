<template>
	<v-app-bar :height="navbarHeight" hide-on-scroll elevate-on-scroll class="navigation" :fixed="this.$route.name !== 'thankyou'" color="rgba(255, 255, 255, .79)">
		<nuxt-link class="navigation__site-title" to="/">
			<img :src="scrolledDown ? $store.state.header.alt_logo.url : $store.state.header.logo.url" :alt="$store.state.header.logo.alt" />
		</nuxt-link>
		<ul class="navigation__items">
			<li v-for="(menuLink, index) in $store.state.header.menu_links" :key="index" class="navigation__menu-item">
				<prismic-link :field="menuLink.link">{{ $prismic.asText(menuLink.label) }}</prismic-link>
			</li>
			<li class="navigation__menu-item">
				<a v-if="$store.state.auth.user == null" @click="$store.dispatch('auth/showAuthModal')" role="button">Login</a>
				<a class="logout-button" @click="signOut" role="button" v-else>Logout</a>				
			</li>
			<li class="navigation__menu-item--user">
				<a href="/user/projects">{{$store.state.auth && $store.state.auth.user ? $store.state.auth.user.email : null}}</a>
			</li>
		</ul>
	</v-app-bar>
</template>
<script>
import Cookie from 'js-cookie'
import { createNamespacedHelpers } from "vuex";
const { mapState, mapActions } = createNamespacedHelpers("auth");
import $debounce from "lodash.debounce"
import { auth } from '~/plugins/firebase.js';
export default {
	name: "AppNavigation",
	components: {},
  
	data() {
		return {
			scrolledDown: false,
			menuContent: [],
			menuLinks: [],
			hasClass: false,
			dialog: false
		}
	},
	mounted() {
		window.addEventListener("scroll", this.isScrolling)
	},
	computed: {
		...mapState({
			visibility: (state) => state.authModal.visibility,
			user: (state) => state.user
    }),
		navbarHeight() {
			var viewportWidth = this.$vuetify.breakpoint.width

			switch (true) {
				case viewportWidth <= 500:
					return "65px"

				case viewportWidth <= 768:
					return "60px"

				case viewportWidth < 960:
					return "70px"

				default:
					return "95px"
			}
		},
	},
	watch: {
    visibility(newVal) {
      newVal ? console.log('true') : console.log('false')
    }
  },
	methods: {
		...mapActions(['hideAuthModal']),
		isScrolling: $debounce(function() {
			this.scrolledDown = window.scrollY > 50 || this.hasClass
		}, 100),
		async signOut() {
			await auth.signOut()
			await Cookie.remove('vuex');
			//this.$router.push('/signedout')
			this.$store.commit('auth/setUser', null)
			window.location.reload()
		}
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
		padding: 4px 0px !important;
		justify-content: space-between;

		@include respond(mobileSmall) {
			flex-direction: column;
		}

		@include respond(mobileSmallLand) {
			flex-direction: row;
		}

		@include respond(mobileLarge) {
			flex-direction: row;
			padding: 4px 40px !important;
		}
	}
}
.navigation {
	@include respond(mobilePort) {
		height: 79px;
	}
	&__site-title {
		color: inherit;
		height: 58px;
		max-width: 194px;
		display: block;

		@include respond(mobilePort) {
			padding: 7px 0;
		}

		@include respond(tabletLarge) {
			max-width:307px;
		}

		@include respond(desktopSmall) {
			max-width:500px;
		}

		img {
			object-fit:contain;
		}
	}

	&__items {
		//flex: 0 0 40%;
		max-width: 1170px;
		justify-content: space-between;
		flex-wrap: wrap;
		display: flex;

		@include respond(mobileLargeMax) {
			background-color:rgba($color-white, .84);
		}
	}

	&__menu-item {
		min-width: 135px;
		height: 30px;
		display: flex;
		justify-content: center;
		@include respond(mobileLarge) {
			height:40px;
		}

		a {
			padding: 0 16px;
			display: flex;
			align-items: center;
			font-weight: 700;
			letter-spacing: 0.15rem;
			text-transform: uppercase;
			position: relative;
			font-size: .8em;
			color: inherit !important;
			text-decoration:none;

			@include respond(mobileLarge) {
				font-size:.9em;
			}

			// @include respond(mobileSmallLand) {
			// 	font-size: 14px;
			// }

			// @include respond(mobileLarge) {
			// 	font-size: 14px;
			// }

			&::before {
				content: "";
				position: absolute;
				left: 0;
				bottom: 2px;
				right: 0;
				opacity: .5;
				width:0;
				height:8px;
				transition: all 0.3s ease-in;
			}
			
			@media(hover) {
				&:hover {
					&::before {
						content:'';
						width:112%;
						background-color:$primary;
						transition:all 0.3s ease-in;
					}
				}				
			}
		}

		&--user {
			align-items:center;
			padding:0 16px;
			display:flex;
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
