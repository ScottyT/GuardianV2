<template>
  <div class="login-form">
    <div class="login-form-box">
      <h1>Login</h1>
      <form ref="loginForm" class="form" method="post">
        <div class="form__form-group">
          <label for="email" class="form__label">Email</label>
          <input type="text" class="Form__input" name="email" id="email" v-model="email" />
        </div>
      </form>
    </div>
  </div>
</template>
<script>
export default {
  name: 'Login',
  props: ['slice']
}
</script>