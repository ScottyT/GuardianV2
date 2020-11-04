<template>
  <div class="user-area">
    <div class="user-area__project">
      <p>{{item ? item.name : null}}</p>
    </div>
  </div>
</template>
<script>
export default {
  async asyncData({ $http, params, error, store }) {
    const id = params.uid
    //const project = await store.dispatch("project/fetchProject", id)
    const project = await store.getters["project/getProject"]
    //console.log(project)
    return { 
      id: id,
      data: project
    }
  },
  computed: {
    item() {
      return this.$store.getters["project/getProject"]
    }
  },
  methods: {
    getProject() {
      return this.$store.dispatch("project/fetchProject", this.id)
    }
  },
  created() {
    this.getProject()
  }
}
</script>
<style lang="scss">
.user-area {
  &__project {
    max-width:90rem;
    width:100%;
    margin:auto;
    padding:20px 4vw;
  }
}
</style>