<template>
	<div :class="imageClass" v-lazyload>
		<img class="img-wrapper" :data-url="source" :alt="alt" />
	</div>
</template>
<script>
import LazyLoadDirective from '../directives/LazyLoadDirective'
export default {
	name: "LazyImage",
	props: {
		alt: String,
		source: String,
		imageClass: String,
	},
	directives: {
		lazyload: LazyLoadDirective
	}
}
</script>
