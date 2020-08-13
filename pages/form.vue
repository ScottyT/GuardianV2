<template>
	<div class="form-wrapper">
		<form method="POST">
			<label>Name</label><input v-model="this.name" name="name" type="text" value="" />
			<label>Address</label>
			<input v-model="this.address" name="address" type="text" value="" />
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
