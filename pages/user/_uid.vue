<template>
  <div class="user-area">
    <slices-block :slices="slices" />
  </div>
</template>
<script>
import { auth } from '~/plugins/firebase.js'
export default {
  middleware({ store, redirect }) {
    if (store.state.auth.user == null) {
      return redirect('/')
    }
  },
  async asyncData({ $prismic, params, error }) {
    const userPages = (await $prismic.api.getByUID("user", params.uid)).data;
    console.log(userPages)
    if (userPages) {
      return {
        slices: userPages.body
      }
    } else {
      error({ statusCode: 404, message: 'Page not found' })
    }
  },
}
</script>
<style lang="scss">
.user-area {
  padding-top:100px;
}
</style>