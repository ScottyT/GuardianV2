<template>
	<div :class="imageClass" v-lazyload>
		<img class="img-wrapper" :src="source" :alt="alt" />
	</div>
</template>
<script>
export default {
	name: "LazyImage",
	props: {
		alt: String,
		source: String,
		imageClass: String,
	},
}
</script>
