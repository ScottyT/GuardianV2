<template>
	<v-app-bar :clipped-left="true" :height="navbarHeight" hide-on-scroll elevate-on-scroll class="navigation" :fixed="this.$route.name !== 'thankyou'" color="rgba(255, 255, 255, .79)">
		<nuxt-link class="navigation__site-title" to="/">
			<img :src="scrolledDown ? $store.state.header.alt_logo.url : $store.state.header.logo.url" :alt="$store.state.header.logo.alt" />
		</nuxt-link>
		<button type="button" class="navigation__toggle-button" @click="menuOpened = !menuOpened">
			<span class="navigation__button-lines--1 navigation__button-lines"></span>
			<span class="navigation__button-lines--2 navigation__button-lines"></span>
			<span class="navigation__button-lines--3 navigation__button-lines"></span>
		</button>
		<ul :class="`navigation__items ${menuOpened ? 'open' : ''}`">
			<li v-for="(menuLink, index) in $store.state.header.menu_links" :key="index" class="navigation__menu-item">
				<prismic-link :field="menuLink.link">{{ $prismic.asText(menuLink.label) }}</prismic-link>
			</li>
			<li class="navigation__menu-item">
				<a v-if="$store.state.auth.user == null" @click="$store.dispatch('auth/showAuthModal')" role="button">Login</a>
				<a class="logout-button" @click="signOut" role="button" v-else>Logout</a>				
			</li>
			<li class="navigation__menu-item--user">
				<a href="/dashboard/projects">{{$store.state.auth && $store.state.auth.user ? $store.state.auth.user.email : null}}</a>
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
			dialog: false,
			menuOpened: false
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
	// watch: {
  //   visibility(newVal) {
  //     newVal ? console.log('true') : console.log('false')
  //   }
  // },
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
			if (this.$route.name == "dashboard-projects-uid") {
				location.href = "/"
			} else {
				window.location.reload()
			}
		}
	}
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
		justify-content: space-around;
		flex-wrap:wrap;

		@include respond(mobileSmall) {
			flex-direction: row;
		}

		@include respond(mobileLarge) {
			flex-direction: row;
			flex-wrap:nowrap;
			justify-content:space-between;
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
		position:relative;
		z-index:-1;

		@include respond(mobileLargeMax) {
			background-color:rgba($color-white, .84);
			transform:translateY(-168px);
			opacity:0;
			transition:all .3s ease-in;
			&.open {
				transition:all .3s ease-in;
				opacity:1;
				transform:translateY(0);
			}
		}
	}

	&__toggle-button {
		width:50px;
		height:40px;
		border:1px solid $color-black;
		display:flex;
		flex-direction:column;
		position:relative;
		justify-content: space-around;
		padding:5px 8px;
		border-radius:5px;
		@include respond(mobileLarge) {
			display:none;
		}
	}
	&__button-lines {
		background-color:$color-black;
		width:100%;
		height:2px;
	//	position:absolute;
		&--1 {
			//top:5px;
		}
		&--2 {
			//top:15px;
		}
		&--3 {
			//top:35px;
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

			a {
				color:$primary!important;
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
