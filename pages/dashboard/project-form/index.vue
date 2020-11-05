<template>
  <div>
    <h1>Create Project Here</h1>
    <ValidationObserver ref="create" v-slot="{ handleSubmit }">
      <form class="form form--project" method="POST" @submit.prevent="handleSubmit()">
        
      </form>
    </ValidationObserver>
  </div>
</template>
<script>
import {
    ValidationProvider,
    ValidationObserver
  } from 'vee-validate';
export default {
  layout: 'dashboard',
  components: {
    ValidationProvider,
    ValidationObserver
  }
}
</script>