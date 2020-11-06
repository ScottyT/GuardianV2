<template>
  <div class="user-area">
    <h1>Create Project Here</h1>
    <ValidationObserver ref="create" v-slot="{ handleSubmit }">
      <form class="form form--project project-create-form" method="POST" @submit.prevent="handleSubmit()">
        <ValidationProvider v-slot="{errors}" vid="name" name="Name" class="form__rounded-input-wrapper">
          <input class="form__rounded-input" placeholder="Project name" v-model="name" type="text" name="name" />
          <span class="form__input--error" v-if="errors.length > 0">{{ errors[0] }}</span>
        </ValidationProvider>
        <span class="form__rounded-input-wrapper form__rounded-input-wrapper--searchbox">
          <autocomplete cssClass="form__rounded-input" placeholderText="Search for clients..." :items="clients" @set_result="selectedClient" />
        </span>
        <ValidationProvider v-slot="{errors}" vid="type" name="Type" class="form__rounded-input-wrapper">
          <!-- <input type="text" v-model="type" placeholder="Project type" class="form__rounded-input" name="type" /> -->
          <autocomplete cssClass="form__rounded-input" placeholderText="Search for category..." :items="['Roofing', 'Gutters','Siding','Flooring','Drywall','Painting']" @set_result="selectedCat" />
          <span class="form__input--error" v-if="errors.length > 0">{{ errors[0] }}</span>
        </ValidationProvider>
        <button type="submit" :class="`button button--submit ${submitting ? 'submitting' : ''}`">{{ buttonText }}</button>
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
    },
    data: () => ({
      submitting: false,
      name: '',
      client: '',
      type: '',
      hidden: true
    }),
    watch: {
      search(value) {
        this.$emit('onchange', value)
        if (value != "") {
          this.hidden = !this.hidden
        }
      }
    },
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
      },
      clients() {
        return this.$store.state.users.map((user) => {
          return user.email
        })
      },
    },
    methods: {
      selectedClient(value) {
        this.client = value
      },
      selectedCat(value) {
        this.type = value
      }
    }
  }
</script>
<style lang="scss">
  .project-form {
    &__suggested-terms {
      

      &.hidden {
        display:none;
      }
    }
  }
</style>