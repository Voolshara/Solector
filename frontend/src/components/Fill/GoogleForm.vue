<template>
  <div>
    <div style="width: available; height: 30px"></div>
    <iframe
        src="https://docs.google.com/forms/d/e/1FAIpQLScNB6N-9y-PxeGAHACTGptyMFm3JLItx-R0oTihAXna4Qc49Q/viewform?embedded=true"
        width="640" height="2237" frameborder="0" marginheight="0" marginwidth="0">Загрузка…
    </iframe>
  </div>
</template>

<script>
export default {
  metaInfo: {
    titleTemplate: '%s | Feed Back',
    meta: [
      {
        name: 'description',
        content: 'Это форма обратной связи. Если вы хотите улучшить наш продукт, или у вас есть идеи для его развития, ответьте на вопросы'
      },
    ],
    link: [
      { rel: 'canonical', href: '/feedback' }
    ]
  },
  name: "GoogleForm"
}
</script>

<style>

</style>