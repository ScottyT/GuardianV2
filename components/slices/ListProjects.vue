<template>
  <div class="projects-list-wrapper">
    <project-form v-if="role == 'admin'" />
    <div class="projects-list">
      <div class="projects-list__project" v-for="(item, i) in $store.state.project.projects" :key="i">
        <p>Project-{{i}}</p>
      </div>
    </div>
  </div>
</template>
<script>
  import {
    fireDb
  } from '~/plugins/firebase.js'
  import {
    auth
  } from '~/plugins/firebase.js'
  export default {
    name: 'ListProjects',
    props: ['slice'],
    data() {
      return {
        projects: [],
        project: {},
        authUser: null,
        role: null
      }
    },
    mounted() {
     // this.$store.dispatch('project/setMessage');
    },
    methods: {
      async getProjects(userData) {
        var projectsRef
        //const projects = []
        const users = []
        var role = this.$store.state.auth.user.role
        this.role = role
        console.log(role)
        if (role == 'admin') {
          await fireDb.collection('projects').get().then((qs) => {
            qs.forEach((doc) => {
              this.projects.push(doc.data())
            })
          })
        } else {
          await fireDb.collection('projects').where("client", "==", userData.uid).get().then((qs) => {
            qs.forEach((doc) => {
              this.projects.push(doc.data())
            })
          })
        }


        // await fireDb.collection('projects').where("client", "==", userData.uid)
        // .get().then((querySnapshot) => {
        //   querySnapshot.forEach((doc) => {
        //     this.projects.push(doc.data())
        //     this.project = doc.data()
        //   })
        // })
      },
      async getAllProjects() {

      }
    },
    created() {
      auth.onAuthStateChanged((user) => {
        if (user) {
          //this.getProjects(user)
        }
      })
    }
  }
</script>
<style lang="scss">
.projects-list-wrapper {
  max-width:120rem;
  width:100%;
  margin:auto;
}
  .projects-list {
    margin: auto;
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    grid-template-rows: repeat(auto-fill, minmax(150px, 1fr));
    column-gap: 20px;
    row-gap: 20px;

    &__project {
      font-size: 1.4em;
    }
  }
</style>