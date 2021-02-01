<template>
  <div>
    <div style="width: available; height: 200px"></div>
    <div class="Wait">Совсем скоро здесь появится инструкция</div>
    <div class="Wait">Ждите</div>
    <div style="width: available; height: 300px"></div>
  </div>
</template>

<script>
export default {
  metaInfo: {
    titleTemplate: '%s | Instruction',
    meta: [
      {
        name: 'description',
        content: 'Здесь вы найдёте подробную пошаговую инструкцию по сборке солнечной электростанции своими руками.'
      },
    ],
    link: [
      { rel: 'canonical', href: '/instruction' }
    ]
  },
  name: "instruction"
}
</script>

<style scoped>
.Wait {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  font-size: 8vmin;
}
</style>