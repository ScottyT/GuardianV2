<template>
	<div class="form-wrapper">
		<form method="POST" @submit.prevent="submitForm">
			<label>Name</label><input v-model="name" name="name" type="text" value="" />
			<label>Address</label>
			<input v-model="address" name="address" type="text" value="" />
			<input type="submit" value="submit" />
		</form>
	</div>
</template>
<script>
export default {
	name: "Form",
	data() {
		return {
			name: "",
			address: "",
			message: "",
		}
	},
	methods: {
		async submitForm() {
			await this.$axios
				.$post("sendform", {
					name: this.name,
					address: this.address,
				})
				.then((res) => {
					this.message = res
				})
				.catch((error) => {
					console.log(error)
				})
		},
	},
}
</script>
<style lang="scss">
.form-wrapper {
	margin-top: 100px;
}
</style>
