<template>
  <div class="user-area">
    <div v-for="(item, i) in data" :key="i">
      {{item.name}}
    </div>
  </div>
</template>
<script>
export default {
  layout: 'dashboard',
  async asyncData({ store, params }) {
    //const user = await store.getters['auth/getUser']
    const projects = await store.state.project.favorites
    return {
      data: projects
    }
  },
  head() {
    return {
      title: 'Favorites'
    }
  }
}
</script>