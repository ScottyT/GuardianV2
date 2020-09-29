<template>
	<v-footer class="justify-center py-4 grey darken-4" dark padless>
		<div class="footer">
			<ul class="footer__items">
				<li v-for="(menuLink, index) in $store.state.footer.menu_links" :key="index" class="footer__menu-item">
					<prismic-link :field="menuLink.link">{{ $prismic.asText(menuLink.label) }}</prismic-link>
				</li>
			</ul>
		</div>
	</v-footer>
</template>
<script>
export default {
	name: "FooterNav",
}
</script>
<style lang="scss">
.footer {
	margin: auto;
	max-width: 114rem;
	padding: 1.7rem 0 2rem 0;

	#services-list {
		padding-left: 15px;
		padding-top: 7px;
	}

	&__site-title {
		width: 100%;
		text-align: center;
		margin-bottom: 1.5rem;
	}

	&__items {
		flex-direction: row;
		justify-content: center;
		width: 100%;		
		padding-left: 0 !important;

		@include respond(mobileLarge) {
			display: flex;
		}
	}

	&__menu-item {
		padding: 0 4px;
		&:not(:last-child) {
			border-right: 1px solid $color-white;
		}

		a {
			display: flex;
			padding: 0 20px;
			background-image: linear-gradient(120deg, transparent 0%, transparent 50%, rgba($color-white, 0.2) 50%);
			background-size: 245%;
			transition: all 0.4s ease-in;
			color: inherit !important;

			&:hover {
				background-position: 100%;
			}
		}

		&--services {
		}
	}
}
</style>
