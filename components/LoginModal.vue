<template>
  <transition name="fadeIn">
  <div class="login-modal modal-backdrop" v-if="$store.state.auth.authModal.visibility">
    <div class="login-form-box">
      <div class="login-form-box__header">
        <h4>Login</h4>
        
        <a role="button" @click="$store.dispatch('auth/hideAuthModal')">close</a>
      </div>
      <div class="alert alert--danger" v-if="getError">{{ $store.state.auth.error }}</div>
      <ValidationObserver v-slot="{ handleSubmit }" tag="form" slim>
      <form ref="loginForm" class="form" method="post" @submit.prevent="handleSubmit(signin)">
        <div class="form__form-group">
          <ValidationProvider v-slot="{ errors }" name="email" rules="required|email" class="form__input--input-group">
            <label for="email" class="form__label">Email</label>
            <input type="text" class="form__input" name="email" id="email" v-model="email" />
            <br/>
            <span class="form__input--error">{{errors[0]}}</span>
          </ValidationProvider>
          <ValidationProvider v-slot="{errors}" name="pass" rules="required" class="form__input--input-group">
            <label for="pass" class="form__label">Password</label>
            <input type="password" class="form__input" id="pass" name="pass" v-model="password" />
            <br/>
            <span class="form__input--error">{{errors[0]}}</span>
          </ValidationProvider>
        </div>
        <a class="small-text color-red">Forgot your password?</a><br/>
        <v-btn type="submit">Login</v-btn>
      </form>
      </ValidationObserver>
      
    </div>
  </div>
  </transition>
</template>
<script>
import { createNamespacedHelpers } from "vuex";
import { ValidationProvider, ValidationObserver } from 'vee-validate';

const { mapActions, mapGetters } = createNamespacedHelpers("auth");
export default {
  name: 'LoginModal',
  components: {
    ValidationProvider,
    ValidationObserver
  },
  data() {
    return {
      email: '',
      password: ''
    }
  },
  computed: {
    ...mapGetters(['getError'])
  },
  methods: {
    ...mapActions(["login", "hideAuthModal"]),
    signin() {     
      this.$store.dispatch('auth/login', {
        email: this.email,
        password: this.password
      }).then(() => {
        this.$refs.loginForm.reset()
        this.email = ''
        this.password = ''
        
      });
    }
  }
}
</script>
<style lang="scss">
.login-form-box {
  background-color:$color-white;
  max-width:500px;
  width:100%;
  padding:20px;
  border:2px solid $color-black;
  border-radius:10px;

  &__header {
    display:flex;
    justify-content:space-between;
  }
}
.alert {
	padding:10px 15px;

	&--danger {
		background-color:#ffdcd6;
		border:2px solid #d60000;
		color:#880707;
	}
}
</style>