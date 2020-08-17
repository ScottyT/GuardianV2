<template>
  <div class="form-wrapper">
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
      <textarea />
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
