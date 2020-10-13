<template>
  <div class="projects-list">
    <div class="projects-list__project" v-for="(item, i) in projects" :key="i">
      <p>Project-{{i}}</p>
    </div>
  </div>
</template>
<script>
import { fireDb } from '~/plugins/firebase.js'
import { auth } from '~/plugins/firebase.js'
export default {
  name: 'ListProjects',
  props: ['slice'],
  data() {
    return {
      projects: [],
      project: {},
      authUser: null
    }
  },
  methods: {
    async getProjects(uid) {
      await fireDb.collection('projects').where("client", "==", uid)
      .get().then((querySnapshot) => {
        querySnapshot.forEach((doc) => {
          this.projects.push(doc.data())
          this.project = doc.data()
        })
      })
    },
    async getAllProjects() {
      await this.$axios.$get('get-projects')
    }
  },
  created() {
    auth.onAuthStateChanged((user) => {
      if (user) {
        this.getProjects(user.uid)
      }
    })
  }
}
</script>
<style lang="scss">
.projects-list {
  max-width:120rem;
  width:100%;
  margin:auto;
  display:grid;
  grid-template-columns:repeat(auto-fill, minmax(300px, 1fr));
  column-gap:20px;
  row-gap:20px;

  &__project {
    font-size:1.4em;
  }
}
</style>