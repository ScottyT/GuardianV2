<template>
  <div class="user-area">
    <div class="user-area__project">
      <p>{{item ? item.name : null}}</p>
    </div>
  </div>
</template>
<script>
export default {
  layout: 'dashboard',
  async asyncData({ $http, params, error, store }) {
    const id = params.uid
  
    const user = await store.getters['auth/getUser']
    const project = await store.getters["project/getProject"]
    //console.log(project)
    return {
      id: id,
      data: project,
      user: user
    }
  },
  async middleware({ params, store, redirect }) {
    
    if (store.state.auth.user == null) {
      return redirect('/')
    }
    await store.dispatch("project/fetchProject", params.uid)
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