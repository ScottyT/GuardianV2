<template>
  <div class="user-area">
    <h1>Create Project Here</h1>
    <h2>{{message}}</h2>
    <ValidationObserver ref="create" v-slot="{ handleSubmit }">
      <form class="form form--project project-create-form" method="POST" @submit.prevent="handleSubmit(create)">
        <ValidationProvider  mode="eager" vid="project" name="project" class="form__rounded-input-wrapper">
          <input class="form__rounded-input" placeholder="Project name" v-model="name" type="text" name="project" />
          <span class="form__input--error">{{ errorsList ? errorsList.Project : null }}</span>
        </ValidationProvider>
        <ValidationProvider v-slot="{errors}" mode="eager" vid="client" name="Client" class="form__rounded-input-wrapper form__rounded-input-wrapper--searchbox">
          <autocomplete cssClass="form__rounded-input" placeholderText="Search for clients..." :items="clients" @set_result="selectedClient" />
          <span class="form__input--error">{{ errors[0] }}</span>
        </ValidationProvider>
        <ValidationProvider v-slot="{errors}" vid="type" mode="eager" name="Type" class="form__rounded-input-wrapper">
          <!-- <input type="text" v-model="type" placeholder="Project type" class="form__rounded-input" name="type" /> -->
          <autocomplete cssClass="form__rounded-input" placeholderText="Search for category..." :items="['Roofing', 'Gutters','Siding','Flooring','Drywall','Painting']" @set_result="selectedCat" />
          <span class="form__input--error">{{ errors[0] }}</span>
        </ValidationProvider>
        <ValidationProvider v-slot="{errors}" vid="date" class="form__rounded-input-wrapper">
          <date-range-picker class="form__rounded-input" ref="picker" :opens="'center'" :time-picker="true" :time-picker24-hour="false" v-model="dateRange">
            <template v-slot:input="picker" style="min-width: 350px">
              <div tabindex="0">{{ picker.startDate | date }} - {{picker.endDate | date }}</div>
            </template>
          </date-range-picker>
          <span class="form__input--error">{{ errors[0] }}</span>
        </ValidationProvider>
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
      errorsList: {},
      hidden: true,
      message: '',
      colors: [{
          key: 'blue',
          value: 'Roofing'
        },
        {
          key: 'indigo',
          value: 'Gutters'
        },
        {
          key: 'deep-purple',
          value: 'Siding'
        },
        {
          key: 'cyan',
          value: 'Flooring'
        },
        {
          key: 'green',
          value: 'Drywall'
        },
        {
          key: 'orange',
          value: 'Painting'
        }
      ],
      color: ''
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
        this.errorsList = {}
        this.submitting = true
        await this.$axios.$post("/sendform", {
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
              type: res.Errors.Type,
              date: res.Errors.DateRange
            })
            
            return;
          }
          
          this.message = res
          this.submitting = false
          this.submitted = true
        })
        .catch((err) => {
          this.message = err.message
        })
        
        for (var x = 0; x < this.colors.length; x++) {
          if (this.type == this.colors[x].value) {
            this.color = this.colors[x].type
          }
        }

        if (Object.keys(this.errorsList).length === 0 && this.errorsList.constructor === Object) {
          fireDb.collection("projects").add({
            client: this.client,
            color: this.color,
            description: this.description,
            name: this.name,
            type: this.type,
            start: this.dateRange.startDate,
            end: this.dateRange.endDate
          }).then((docRef) => {
            this.message = "Project was added successfully!"
            this.submitting = false
            this.submitted = true
            this.project = ""
            this.client = ""
            this.type =""
            this.dateRange = {
              startDate: null,
              endDate: null
            }
            this.description = ""
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