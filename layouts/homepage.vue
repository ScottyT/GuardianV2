<template>
	<v-app>
		<app-navigation />
		<main class="v-content__wrap homepage">
			<nuxt />
		</main>
		<footer-nav />
		<login-modal @success="success" />
	</v-app>
</template>
<script>
import { createNamespacedHelpers } from "vuex";
const { mapActions } = createNamespacedHelpers("auth");
export default {
	name: "homepage",
	head() {
		return {
			title: "Home",
		}
	},
	async middleware({ store, $prismic }) {
		await store.dispatch("fetchMenu", $prismic)
	},
	methods: {
		...mapActions(['hideAuthModal'])
	}
}
</script>
<style lang="scss"></style>
