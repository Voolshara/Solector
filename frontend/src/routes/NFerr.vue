<template>
  <div>
    <div style="width: available; height: 200px"></div>
    <div class="notF">
      <h1>404</h1>
      <h2>Похоже, вы ошиблись адресом :)</h2>
    </div>
    <div style="width: available; height: 300px"></div>
  </div>
</template>

<script>
export default {
  metaInfo: {
    titleTemplate: '%s | 404',
    meta: [
      {
        name: 'description',
        content: 'Похоже, вы ошиблись адресом, проверьте ссылку'
      },
    ]
  },
  name: "NFerr"
}
</script>

<style>
.notF {
  width: available;
  text-align: center;
  height: 100%;
}

@media (max-width: 480px) {
  .notF {
    position: relative;
    top: 50%;
    height: initial;
    -webkit-transform: translateY(-50%);
    -ms-transform: translateY(-50%);
    transform: translateY(-50%);
  }
}

.notF h1 {
  margin: 0;
  font-size: 130px;
  font-weight: 300;
}

@media (min-width: 480px) {
  .notF h1 {
    position: relative;
    top: 50%;

  }
}

@media (min-width: 768px) {
  .notF h1 {
    font-size: 220px;
  }
}

</style>