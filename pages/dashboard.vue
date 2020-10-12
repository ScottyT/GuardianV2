<template>
  <h1>Welcome, <span v-if="$store.state.auth.user != null">{{$store.state.auth.user.name}}</span></h1>
</template>
<script>
export default {
  name: 'Dashboard',

}
</script>