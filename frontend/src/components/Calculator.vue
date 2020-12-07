<template>
  <div>
    <div class="topnav">
      <a href=# class="selected">ПО ЭЛЕКТРОПРИБОРАМ</a>
      <a href=# >ПО ЭНЕРГОПОТРЕБЛЕНИЮ</a>
      <a id="menu" href="#" class="icon">&#9776;</a>
    </div>
  <Form/>
  </div>
</template>

<script>
  import Form from '@/components/Calc_forms/Calc-for-device'
  export default {
    components:{
      Form,
    }
  }
</script>

<style>
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
  //font-style: 14px;
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
</style>