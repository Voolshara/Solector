<template>
  <table>
    <tr>
      <td>
        <h2>Рекомендуемая панель</h2>
      </td>
      <td>
        <h2>Выработка</h2>
      </td>
    </tr>
    <tr>
      <td><a v-bind:href="response_data['panel_link']" target="_blank"><img v-bind:src="response_data['img_link']" alt="Sorry we can't upload image"
                          style="height: 400px; width: 350px"></a></td>
      <td>
        <table>
          <tr>
            <td>Модель:</td>
            <td>{{response_data['name']}}</td>
          </tr>
          <tr>
            <td>Производитель:</td>
            <td>{{response_data['company']}}</td>
          </tr>
          <tr>
            <td>Колличество панелей:</td>
            <td>{{response_data['kol']}}</td>
          </tr>
          <tr>
            <td>Общая стоимость панелей:</td>
            <td>{{response_data['cost']}} руб</td>
          </tr>
          <tr>
            <td>Колличество панелей:</td>
            <td>{{response_data['kol']}}</td>
          </tr>
        </table>
      </td>
      <td>
        <graph v-bind:chart_data="response_data" class="graph"/>
      </td>
    </tr>
  </table>
</template>

<script>
import Graph from "@/components/Calc_forms/graph";

export default {
  components: {
    Graph,
  },
  props: {
    response_data: {
      type: Object,
      default: null
    },
  },
  name: "form-response",
}
</script>

<style scoped>
.graph {
  height: 400px;
  width: 600px;
}
</style>