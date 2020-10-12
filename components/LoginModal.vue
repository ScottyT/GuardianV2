<template>
  <transition name="fadeIn">
    <div class="login-modal modal-backdrop" v-if="$store.state.auth.authModal.visibility">
      <div class="login-form-box">
        <div class="login-form-box__header">
          <h4>{{ !forgotPassword ? 'Login' : 'Password Reset' }}</h4>
          <div class="login-form-box__close" @click="closeModal">
            <font-awesome-icon size="2x" :icon="['fas', 'times']" />
          </div>
        </div>
        <div class="login-wrapper" v-if="!forgotPassword">
          <div class="alert alert--danger" v-if="getError">{{ $store.state.auth.error }}</div>
          <ValidationObserver v-slot="{ handleSubmit }" tag="form" slim>
            <form ref="loginForm" class="form" method="post" @submit.prevent="handleSubmit(signin)">
              <div class="form__form-group">
                <ValidationProvider v-slot="{ errors }" name="email" rules="required|email"
                  class="form__input--input-group">
                  <label for="email" class="form__label">Email</label>
                  <input type="text" class="form__input" name="email" id="email" v-model="email" />
                  <br />
                  <span class="form__input--error">{{errors[0]}}</span>
                </ValidationProvider>
                <ValidationProvider v-slot="{errors}" name="pass" rules="required" class="form__input--input-group">
                  <label for="pass" class="form__label">Password</label>
                  <input type="password" class="form__input" id="pass" name="pass" v-model="password" />
                  <br />
                  <span class="form__input--error">{{errors[0]}}</span>
                </ValidationProvider>
              </div>
              <a @click="passwordReset" class="small-text color-red">Forgot your password?</a><br />
              <v-btn type="submit">{{ loggingIn ? 'Logging in' : 'Login' }}</v-btn>
            </form>
          </ValidationObserver>
        </div>
        <div class="forgot-password-wrapper" v-else-if="forgotPassword">
          <p v-if="errors == null">{{emailSentMessage}}</p>
          <div class="alert alert--danger" v-if="getError">{{ $store.state.auth.error }}</div>
          <div class="form__form-group">
            <label class="form__label">Enter your email</label>
            <input type="text" class="form__input" id="email" name="email" v-model="passwordResetEmail" />
            <a v-show="passwordResetEmail !== ''" @click="sendPasswordReset" class="button button--red">{{sendingEmail ? 'Sending...' : 'Send Password Reset'}}</a>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>
<script>
import '@fortawesome/fontawesome-svg-core/styles.css'
  import {
    createNamespacedHelpers
  } from "vuex";
  import {
    ValidationProvider,
    ValidationObserver
  } from 'vee-validate';
  import { auth } from '~/plugins/firebase.js'

  const {
    mapActions,
    mapGetters
  } = createNamespacedHelpers("auth");
  export default {
    name: 'LoginModal',
    components: {
      ValidationProvider,
      ValidationObserver
    },
    data() {
      return {
        email: '',
        password: '',
        forgotPassword: false,
        passwordResetEmail: '',
        emailSentMessage: '',
        loggingIn: false,
        sendingEmail: false
      }
    },
    computed: {
      ...mapGetters(['getError']),
      error: {
        get: function() {
          return this.getError
        },
        set: function(newValue) {
          return null
        }
      }
    },
    methods: {
      ...mapActions(["login", "hideAuthModal"]),
      signin() {
        this.loggingIn = true
        this.$store.dispatch('auth/login', {
          email: this.email,
          password: this.password
        }).then(() => {
          this.$refs.loginForm.reset()
          this.email = ''
          this.password = ''
          this.loggingIn = false
          //this.closeModal()
        });
      },
      closeModal() {
        this.$store.dispatch('auth/hideAuthModal');
        this.forgotPassword = false
      },
      async sendPasswordReset() {
        this.sendingEmail = true
        this.$store.dispatch('auth/passwordReset', this.passwordResetEmail).then(() => {
          this.emailSentMessage = "Reset email is sent"
          this.sendingEmail = false
          this.passwordResetEmail = ''
        })
      },
      passwordReset() {
        this.forgotPassword = !this.forgotPassword
      }
    }
  }
</script>
<style lang="scss">
  .login-form-box {
    background-color: $color-white;
    max-width: 500px;
    width: 100%;
    padding: 20px;
    border: 2px solid $color-black;
    border-radius: 10px;

    &__header {
      display: flex;
      justify-content: space-between;
    }

    &__close {
      padding:2px 10px;
      cursor:pointer;
    }
  }

  .alert {
    padding: 10px 15px;

    &--danger {
      background-color: #ffdcd6;
      border: 2px solid #d60000;
      color: #880707;
    }
  }
</style>