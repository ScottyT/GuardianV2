<template>

  <div class="user-area">
    <h1>Create Project Here</h1>
    <ValidationObserver ref="create" v-slot="{ handleSubmit }">
      <form class="form form--project" method="POST" @submit.prevent="handleSubmit()">
        <ValidationProvider v-slot="{errors}" vid="name" name="Name" class="form__rounded-input-wrapper">
          <input class="form__rounded-input" placeholder="Project name" v-model="name" type="text" name="name" />
          <span class="form__input--error" v-if="errors.length > 0">{{ errors[0] }}</span>
        </ValidationProvider>
        <span class="form__rounded-input-wrapper form__rounded-input-wrapper--searchbox">
          <input type="text" v-model="search" class="form__rounded-input" placeholder="Search for client..." />
          <ul v-if="!hidden" class="project-form__suggested-terms">
            <li v-for="(item, i) in filteredClients" :key="i">{{ item.email }}</li>
          </ul>

        <!-- <select v-model="selected" class="form__rounded-input">
            <option disabled value="">Selected a client</option>
            <option v-for="(client, i) in $store.state.users" :key="i">
              {{client.email}}
            </option>
          </select> -->
        </span>
        <ValidationProvider v-slot="{errors}" vid="type" name="Type" class="form__rounded-input-wrapper">
          <input type="text" v-model="type" placeholder="Project type" class="form__rounded-input" name="type" />
          <span class="form__input--error" v-if="errors.length > 0">{{ errors[0] }}</span>
        </ValidationProvider>
        <button type="submit"
          :class="`button button--submit ${submitting ? 'submitting' : ''}`">{{ buttonText }}</button>
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
      selected: '',
      type: '',
      search: '',
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
        return this.$store.state.users
      },
      filteredClients() {
        let searchText = this.search
        searchText.trim().toLowerCase();
        return this.clients.filter(user => {
          return user.email.toLowerCase().includes(searchText) && searchText != ''
        })
      }
    },
    methods: {
      handler: function(evt) {
        this.$emit("change", evt)
      }
    }
  }
</script>
<style lang="scss">
  .project-form {
    &__suggested-terms {
      z-index:4;
      position:relative;
      width:500px;
      background-color:$color-white;
      padding:15px;
      border:1px solid $color-black;

      &.hidden {
        display:none;
      }
    }
  }
</style>