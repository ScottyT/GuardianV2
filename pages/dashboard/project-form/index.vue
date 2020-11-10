<template>
  <div class="user-area">
    <span>{{message}}</span>
 
    <div>Errors: {{errorsList}}</div>
    <h1>Create Project Here</h1>
    <ValidationObserver ref="create" v-slot="{ handleSubmit }">
      <form class="form form--project project-create-form" method="POST" @submit.prevent="handleSubmit(create)">
        <ValidationProvider v-slot="{errors}" vid="project" name="project" class="form__rounded-input-wrapper">
          <input class="form__rounded-input" placeholder="Project name" v-model="name" type="text" name="project" />
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
        <span class="form__rounded-input-wrapper">
          <date-range-picker ref="picker" :opens="'center'" :time-picker="true" :time-picker24-hour="false" v-model="dateRange">
            <template v-slot:input="picker" style="min-width: 350px">
              {{ picker.startDate | date }} - {{picker.endDate | date }}
            </template>
          </date-range-picker>
        </span>
        <span class="form__rounded-input-wrapper--textarea form__rounded-input-wrapper">
          <textarea class="form__rounded-input" v-model="description" placeholder="Description" name="description"></textarea>
        </span>
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
  import { fireDb } from '~/plugins/firebase.js'
  import DateRangePicker from 'vue2-daterange-picker'
  import 'vue2-daterange-picker/dist/vue2-daterange-picker.css'
  export default {
    layout: 'dashboard',
    components: {
      ValidationProvider,
      ValidationObserver,
      DateRangePicker
    },
    data: () => ({
      submitting: false,
      submitted: false,
      name: '',
      client: '',
      type: '',
      description: '',
      dateRange: {
        startDate: null,
        endDate: null
      },
      errorsList: null,
      hidden: true,
      message: ''
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
      },
      async create() {
        this.errorsList = null
        this.submitting = true
        await this.$axios.$post("/create", {
          project: this.name,
          client: this.client,
          dateRange: this.dateRange,
          description: this.description,
          type: this.type,
          formtype: 'ProjectCreate'
        }).then((res) => {
          if (res.Errors) {
            this.errorsList = res.Errors
            this.$refs.create.setErrors({
              project: res.Errors.ProjectName,
              client: res.Errors.Client,
              type: res.Errors.Type
            })
            this.submitting = false
            this.submitted = false
            return;
          }
          console.log("outside firedb")
          
          this.message = res
          this.submitting = false
          this.submitted = true
        })
        .catch((err) => {
          this.message = err.message
        })

        if (Object.keys(this.errorsList).length === 0 && this.errorsList.constructor === Object) {
          fireDb.collection("projects").add({
            client: this.client,
            description: this.description,
            name: this.name,
            type: this.type,
            start: this.dateRange.startDate,
            end: this.dateRange.endDate
          }).then((docRef) => {
            console.log("Inside fire db:", docRef.id)
            this.message = res.message
            this.submitting = false
            this.submitted = true
          }).catch((error) => {
            console.error(error)
          })
        }
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