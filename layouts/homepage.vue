<template>
	<v-app>
		<app-navigation />
		<main class="v-content__wrap homepage">
			<nuxt />
		</main>
		<footer-nav />
	</v-app>
</template>
<script>
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
}
</script>
<style lang="scss"></style>
