<template>
  <div id="app">
    <Main msg=""/>
    <Footer />
  </div>
</template>

<script lang="ts">
import Main from './components/Main.vue';
import Footer from './components/Footer.vue';

export default {
  components: {
    Main,
    Footer,
  }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
