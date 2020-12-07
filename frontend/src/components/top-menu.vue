<template>
  <div>
    <div class="parallax"></div>
    <div class="topnav" id="myTopnav">
      <a href="/">MARKETPLACE</a>
      <a href=/selection>ПОДБОР</a>
      <a href="instruction">ИНСТРУКЦИЯ</a>
      <a href="http://solector.tilda.ws/" target="_blank">О НАС</a>
    </div>
  </div>
</template>

<script>
export default {
name: "top-menu"
}
</script>

<style>
*{
  margin: 0;
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

@media screen and (max-width: 768px) {
  .topnav a:not(:first-child) {
    display: none;
  }

  nav {
    width: 100px;
    height: 100px;
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