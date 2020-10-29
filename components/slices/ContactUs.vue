<template>
  <div class="form-wrapper">
    <h1 class="form-wrapper__heading red-text small text-center">OUR SHOWROOM IS YOUR RESTORED PROPERTY!</h1>
    <div class="form-wrapper__emergency-box">
      <img :class="broken ? 'hidden' : ''" @click.once="broken = !broken"
        src="https://images.prismic.io/guardianrestoration/fef8fd2f-273c-4935-9fb7-4ab62b4985e5_Emergency+Box_Unbroken.png?auto=compress,format" />
      <div :class="`form-wrapper__call-button ${broken ? '' : 'form-wrapper__call-button--hidden'}`">
        <a href="tel:5555555555" class="click-area"></a>
        <img :class="clicked ? '' : 'form-wrapper__call-button--hidden'"
          src="https://images.prismic.io/guardianrestoration/c97f5bc1-bacc-458f-a781-10282179e38f_Emergency+Box_BROKEN-ButtonPushDown.png?auto=compress,format" />
        <img :class="clicked ? 'form-wrapper__call-button--hidden' : ''"
          src="https://images.prismic.io/guardianrestoration/9290df64-a7d6-4b06-b023-29addd65d0b4_Emergency+Box_BROKEN.png?auto=compress,format" />
      </div>
    </div>
    <div class="form-wrapper__form-box">
      <div class="form-wrapper__form-box-intro">
        <h2>WE'RE LISTENING</h2>
        <p class="form-wrapper__form-box-intro--cta">
          Leave us your name, email, and a brief description of the issue... the restoration you need is a message away.
        </p>
      </div>
      <ValidationObserver ref="form" v-slot="{ handleSubmit }">
        <form class="form-wrapper__form form--contact form" method="POST" @submit.prevent="handleSubmit(submitForm)">
          <h2 v-if="successMessage">{{ successMessage }}</h2>
          <p v-if="errorsList.length">
            <strong>Please correct the follwing errors:</strong>
          </p>
          <ul v-if="errorsList.length">
            <li v-for="error in errors" :key="error.id">{{ error }}</li>
          </ul>
          <ValidationProvider v-slot="{errors}" mode="eager" vid="name" name="Name" class="form__rounded-input-wrapper">
            <span class="form__icon">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
              </svg>
            </span>
            <input class="form__rounded-input form__rounded-input--name" :class="errors.length > 0 ? `is-false` : ''"
              v-model="name" name="name" placeholder="name" type="text" value />
            <span class="form__input--error">{{ errors[0] }}</span>
          </ValidationProvider>
          <ValidationProvider v-slot="{errors}" mode="eager" vid="email" name="Email"
            class="form__rounded-input-wrapper">
            <span class="form__icon">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4" />
              </svg>
            </span>
            <input class="form__rounded-input form__rounded-input--email" :class="errors.length > 0 ? `is-false` : ''"
              v-model="email" name="email" placeholder="email" value />
            <span class="form__input--error">{{ errors[0] }}</span>
          </ValidationProvider>
          <ValidationProvider v-slot="{errors}" mode="eager" vid="message" name="Message"
            class="form__rounded-input-wrapper form__rounded-input-wrapper--textarea">
            <span class="form__icon form__icon--textarea-icon">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
            </span>
            <textarea class="form__rounded-input form__rounded-input--message"
              :class="errors.length > 0 ? `is-false` : ''" placeholder="message" v-model="message" name="message"
              type="text"></textarea>
            <span class="form__input--error">{{ errors[0] }}</span>
          </ValidationProvider>
          <div class="button-bg">
            <img :class="`wing wing--left ${submitted ? 'visible' : ''}`"
              src="https://images.prismic.io/guardianrestoration/3fa63f63-2b10-42a4-a5b1-1651a3c13fa9_Wing+Small+for+Guardian+home+2.png?auto=compress,format" />
            <button :disabled="submitted" type="submit"
              :class="`button button--submit ${submitting ? 'submitting' : ''}`">{{ buttonText }}</button>
            <img :class="`wing wing--right ${submitted ? 'visible' : ''}`"
              src="https://images.prismic.io/guardianrestoration/3fa63f63-2b10-42a4-a5b1-1651a3c13fa9_Wing+Small+for+Guardian+home+2.png?auto=compress,format" />
          </div>
        </form>
      </ValidationObserver>
    </div>
  </div>
</template>
<script>
  import anime from "animejs"
  import {
    ValidationProvider,
    ValidationObserver
  } from 'vee-validate';
  export default {
    name: "ContactUs",
    props: ["slice"],
    components: {
      ValidationProvider,
      ValidationObserver
    },
    data: () => ({
      name: "",
      email: "",
      message: "",
      successMessage: "",
      errorsList: [],
      broken: false,
      clicked: false,
      submitting: false,
      submitted: false
    }),
    computed: {
      buttonText() {
        switch (true) {
          case this.submitting == true:
            return 'SUBMITTING'
            break;
          case this.submitted == true:
            return 'SUBMITTED'
            break;
          default:
            return 'SUBMIT'
        }
      }
    },
    mounted() {
      const buttonArea = document.querySelector('.form-wrapper__call-button .click-area');
      const submitBtn = document.querySelector('.form--contact .button--submit');
      buttonArea.addEventListener('mousedown', () => {
        this.clicked = true
      })
      buttonArea.addEventListener('mouseup', () => {
        this.clicked = false
      })
      buttonArea.addEventListener('touchstart', () => {
        this.clicked = true
      })
      buttonArea.addEventListener('touchcancel', () => {
        this.clicked = false
      })

      submitBtn.addEventListener('mousedown', () => {
        submitBtn.classList.add('pushed')
      })
      submitBtn.addEventListener('mouseup', () => {
        submitBtn.classList.remove('pushed')
      })
    },
    methods: {
      async submitForm() {
        const xMax = 10
        const shake = anime({
          targets: ".form--contact",
          easing: "easeInOutSine",
          duration: 550,
          translateX: [{
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
        //this.submitted = true
        this.submitting = true
        await this.$axios
          .$post("/sendform", {
            name: this.name,
            email: this.email,
            message: this.message,
          })
          .then((res) => {
            if (res.Errors) {
              this.$refs.form.setErrors({
                email: res.Errors.Email,
                name: res.Errors.Name,
                message: res.Errors.Message
              })
              shake.play()
              this.submitting = false
              this.submitted = false
              return;
            }
            this.successMessage = res
            this.submitting = false
            this.submitted = true
            //this.$router.push("/thankyou")
          })
          .catch((error) => {
            console.log(error)
            this.submitted = false
            this.submitting = false
            this.successMessage = error.message
          })

      },
    },
  }
</script>
<style lang="scss">
  .form-wrapper {
    padding: 0 4vw 40px;
    position: relative;
    border: 10px solid $color-white;
    background-size: cover;
    background-image: url('https://images.prismic.io/guardianrestoration/69818d1d-e79a-496e-9a54-8ca3de0c10cc_Concrete+Wall+for+Our+Services+Guardian+3.3+MB.jpg?auto=compress,format');

    &:before {
      content: '';
      position: absolute;
      top: 0;
      width: 0;
      height: 0;
      border-left: 30px solid transparent;
      border-right: 30px solid transparent;
      border-top: 30px solid $color-white;
    }

    &__form-box {
      background-image: url('https://images.prismic.io/guardianrestoration/a142f753-d2bc-4e4f-a859-9bec4cd486f7_Contac+Us+Form+no+shadow.png?auto=compress,format');
      background-size: cover;
      background-repeat: no-repeat;
      max-width: 900px;
      margin: auto;
      //height:900px;
      position: relative;
      box-shadow: -3px 6px 10px 0 rgba(0, 0, 0, .42);
    }

    &__form-box-intro {
      width: 100%;
      //@include absCentered(20%, 50%);
      text-align: center;
      padding-top: 77px;
      padding-bottom: 90px;

      &--cta {
        background: rgba($color-white, .7);
        padding: 20px 8vw;
        box-shadow: 0px 7px 2px 0px rgba(0, 0, 0, .24);
        margin-top: 15px;
      }
    }

    &__emergency-box {
      //max-width:670px;
      height: 100vw;
      margin: auto;
      position: relative;
      margin-bottom: 40px;

      @include respond(mobileLarge) {
        height: 100vh;
        display: flex;
        justify-content: center;
      }

      &>img {
        cursor: pointer;
        width: auto;
      }

      img {
        position: absolute;
        top: 0;
        object-fit: contain;

        &.hidden {
          display: none;
        }
      }
    }

    &__heading {
      max-width: 600px;
      margin: auto;
      font-style: italic;
      padding: 40px 0;
    }

    &__call-button {
      position: absolute;
      top: 0;
      width: 100%;
      height: 100%;

      .click-area {
        cursor: pointer;
        @include absCentered(48%, 50%);
        z-index: 4;
        width: 35vw;
        height: 35vw;

        @include respond(mobileLarge) {
          width: 280px;
          height: 275px;
        }
      }

      &--hidden {
        display: none;
      }
    }

    &__form {
      display: flex;
      max-width: 800px;
      padding: 15px;
      flex-direction: column;
      margin: auto;
    }
  }
</style>