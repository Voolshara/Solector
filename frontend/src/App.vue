<template>
  <div id="app">
    <link rel="shortcut icon" href="public/favicon.ico">
    <Menu/>
    <router-view/>
    <Footer/>
  </div>
</template>

<script>
import Menu from "@/components/top-menu.vue";
import Footer from "@/components/footer";
export default {
  components: {
    Menu,
    Footer,
  }
}

</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
