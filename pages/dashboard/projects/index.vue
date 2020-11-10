<template>

  <div class="user-area">
      <project-calendar :listprojects="projects" :user="data" />
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
import { createNamespacedHelpers } from "vuex";
const { mapState } = createNamespacedHelpers("project");
export default {
  layout: 'dashboard',
  async asyncData({ params, store }) {
    const user = await store.getters['auth/getUser']
    const projects = await store.getters['project/getProjects']
    console.log(store.getters["auth/getUser"])
    console.log(store.getters["project/getProjects"])
    return {
      data: user,
     // projects: projects
    }
  },
  computed: {
    ...mapState({
      projects: (state) => state.projects
    })
  }
}
</script>
<style lang="scss">
.slide-right {
  opacity:0;
  transform:translateX(200px);
  transition: all .2s ease-in;
}
.slide-left {
  opacity:1;
  transform:translateX(-200px)
}
</style>