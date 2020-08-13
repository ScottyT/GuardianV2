<template>
	<div class="form-wrapper">
		<form method="POST" action="/thankyou">
			<label>Name</label><input name="name" type="text" value="" />
			<label>Address</label>
			<input name="address" type="text" value="" />
			<input type="submit" value="submit" />
		</form>
	</div>
</template>
<script>
export default {
	name: "Form",
	data() {
		return {
			Name: "",
			Address: "",
		}
	},
}
</script>
<style lang="scss">
.form-wrapper {
	margin-top: 100px;
}
</style>
