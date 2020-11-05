<template>
  <div class="user-area">
    <project-calendar :user="data" />
  </div>
</template>
<script>
export default {
  layout: 'dashboard',
  async asyncData({ params, store }) {
    const user = await store.getters['auth/getUser']
    console.log(store.getters["auth/getUser"])
    return {data: user}
  }
}
</script>