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
  metaInfo: {
    title: 'Solector',
    titleTemplate: '%s | Main Page',
    htmlAttrs: {
      reptilian: 'gator'
    },
    headAttrs: {
      nest: 'eggs'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' }
    ]
  },
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
