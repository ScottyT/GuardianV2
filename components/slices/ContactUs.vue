<template>
	<div class="form-wrapper">
		<div class="form-wrapper__emergency-box" @click.once="broken = !broken">
			<img :class="broken ? 'hidden' : ''" src="https://images.prismic.io/guardianrestoration/fef8fd2f-273c-4935-9fb7-4ab62b4985e5_Emergency+Box_Unbroken.png?auto=compress,format" />
			<img :class="broken ? '' : 'hidden'" src="https://images.prismic.io/guardianrestoration/9290df64-a7d6-4b06-b023-29addd65d0b4_Emergency+Box_BROKEN.png?auto=compress,format" />
			<div :class="`form-wrapper__call-button ${broken ? '' : 'form-wrapper__call-button--hidden'}`">
				<div class="click-area" @click="clicked = !clicked"></div>
				<img :class="clicked ? '' : 'form-wrapper__call-button--hidden'" src="https://images.prismic.io/guardianrestoration/c97f5bc1-bacc-458f-a781-10282179e38f_Emergency+Box_BROKEN-ButtonPushDown.png?auto=compress,format" />
				<img :class="clicked ? 'form-wrapper__call-button--hidden' : ''" src="https://images.prismic.io/guardianrestoration/9290df64-a7d6-4b06-b023-29addd65d0b4_Emergency+Box_BROKEN.png?auto=compress,format" />
			</div>
		</div>
		<form class="form-wrapper__form form--contact" method="POST" @submit.prevent="submitForm">
			<h2 v-if="this.successMessage">{{ successMessage }}</h2>
			<p v-if="errors.length">
				<strong>Please correct the follwing errors:</strong>
			</p>
			<ul>
				<li v-for="error in errors" :key="error.id">{{ error }}</li>
			</ul>
			<label>Name</label>
			<input class="form__input" v-model="name" name="name" type="text" value />
			<label>Email</label>
			<input class="form__input" v-model="email" name="email" type="email" value />
			<label>Message</label>
			<textarea class="form__input" v-model="message" name="message" type="text" value=""></textarea>
			<v-btn type="submit">Submit</v-btn>
		</form>
	</div>
</template>
<script>
import anime from "animejs"
export default {
	name: "ContactUs",
	props: ["slice"],
	data() {
		return {
			name: "",
			email: "",
			message: "",
			successMessage: "",
			errors: [],
			broken: false,
			clicked: false
		}
	},
	methods: {
		async submitForm() {
			const xMax = 10
			const shake = anime({
				targets: ".form--contact",
				easing: "easeInOutSine",
				duration: 550,
				translateX: [
					{
						value: xMax * -1,
					},
					{
						value: xMax,
					},
					{
						value: xMax / -2,
					},
					{
						value: xMax / 2,
					},
					{
						value: 0,
					},
				],
				autoplay: false,
			})
			const errorVisual = anime({
				targets: ".form--contact",
				borderColor: "#D11726",
				easing: "linear",
				duration: 500,
				delay: 500,
				autoplay: false,
			})
			if (this.name && this.email && this.message) {
				await this.$axios
					.$post("sendform", {
						name: this.name,
						email: this.email,
						message: this.message,
					})
					.then((res) => {
						this.successMessage = res
						this.$router.push("/thankyou")
					})
					.catch((error) => {
						this.errors = error
					})
			}
			this.errors = []
			if (!this.name) this.errors.push("Name required")
			if (!this.email) this.errors.push("Email required")
			if (!this.message) this.errors.push("Message Required")

			if (this.errors) {
				shake.play()
				errorVisual.play()
			}
		},
	},
}
</script>
<style lang="scss">
.form-wrapper {
	padding: 0 4vw;
	background-image:url('https://images.prismic.io/guardianrestoration/69818d1d-e79a-496e-9a54-8ca3de0c10cc_Concrete+Wall+for+Our+Services+Guardian+3.3+MB.jpg?auto=compress,format');

	&__emergency-box {
		max-width:670px;
		height:900px;
		margin:auto;
		position:relative;
		img {
			position: absolute;
			top:0;
			&.hidden {
				display:none;
			}
		}
	}

	&__call-button {
		position:absolute;
		top:0;
		width:100%;
		height:100%;
		.click-area {
			cursor:pointer;
			width:250px;
			height:250px;
			@include absCentered(50%, 50%);
			z-index:4;
		}
		&--hidden {
			display:none;
		}
	}

	&__form {
		display: flex;
		max-width: 800px;
		padding: 15px;
		flex-direction: column;
		margin: auto;
		border: 1px solid $color-black;
		background: $color-white;
		// filter: drop-shadow(2px 3px 3px rgba(0, 0, 0, 0.5));
		box-shadow: 0px 10px 5rem 2px rgba(68, 64, 64, 0.38);
	}
}
</style>
