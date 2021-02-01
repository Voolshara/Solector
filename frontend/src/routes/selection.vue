<template>
  <div id="app" style="width: available">
    <div>
      <calc-solution/>
    </div>
    <div>
    </div>
  </div>
</template>


<script>
import CalcSolution from '@/components/Calculator'

export default {
  name: 'App',
  metaInfo: {
    title: 'Calculator',
    titleTemplate: '%s | Device',
    meta: [
      {
        name: 'description',
        content: 'Solector - это интернет-сервис, где вы можете быстро подобрать комплект солнечных панелей, который будет наиболее выгоден вам'
      },
    ],
    link: [
      { rel: 'canonical', href: '/selection' }
    ]
  },
  components: {
    CalcSolution,
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

* {
  margin: 0;
  padding: 0;
}

a {
  text-decoration: none;
}

li {
  text-decoration: none;
  list-style: none;
}

body {
  font-family: 'Raleway', sans-serif;
}

nav {
  margin: auto;
  width: 800px;
  height: 100px;
}

.topnav {
  background-color: #fff;
  margin-top: 40px;
  text-align: center;
}

.topnav a {
  color: #000;
  text-align: center;
  padding: 14px 16px;
  font-size: 15px;
}

.topnav a:hover {
  border-bottom: 2px solid #808080;
}

.topnav .icon {
  display: none;
}

.topnav .selected {
  border-bottom: 2px solid #000;
}

@media screen and (max-width: 768px) {
  .topnav a:not(:first-child) {
    display: none;
  }

  .topnav a.icon {
    float: right;
    margin-top: -17px;
    display: block;
    color: black;
  }

  nav {
    width: 100px;
    height: 100px;
  }

  .topnav.responsive a.icon {
    position: absolute;
    right: 0;
    top: 0;
  }

  .topnav.responsive a {
    float: none;
    display: block;
    text-align: left;
  }
}

table {
  margin: auto;
}

td {
  padding: 10px 10px;
}

.parallax {
  background-image: url("../assets/good.jpeg");
  min-height: 400px;
  background-attachment: fixed;
  width: available;
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
}

</style>
