<template>
  <div class="search-bar">
    <input type="text" v-model="search" />
  </div>
</template>
<script>
export default {
  name: "SearchBar",
  props: {
    items: {
      type: Array,
      required: false
    }
  },
  data() {
    return {
      search:''
    }
  }
}
</script>