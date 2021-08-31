<template>
  <div id="app">
    <ApiInput />
  </div>
</template>

<script>
import ApiInput from './components/ApiInput.vue'

export default {
  name: 'App',
  components: {
    ApiInput
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin-top: 15vh;
  margin-left: 10vh;
}
</style>
