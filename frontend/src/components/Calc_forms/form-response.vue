<template>
  <div>
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
        <td><a v-bind:href="response_data['panel_link']" target="_blank"><img v-bind:src="response_data['img_link']"
                                                                              alt="Sorry we can't upload image"
                                                                              style="height: 400px; width: 350px"></a>
        </td>
        <td>
          <table>
            <tr>
              <td>Модель:</td>
              <td>{{ response_data['name'] }}</td>
            </tr>
            <tr>
              <td>Производитель:</td>
              <td>{{ response_data['company'] }}</td>
            </tr>
            <tr>
              <td>Количество панелей:</td>
              <td>{{ response_data['kol'] }}</td>
            </tr>
            <tr>
              <td>Общая стоимость панелей:</td>
              <td>{{ response_data['cost'] }} руб</td>
            </tr>
          </table>
        </td>
        <td>
          <graph v-bind:chart_data="response_data" class="graph"/>
        </td>
      </tr>
    </table>
    <br>
    <div style="width: 100px; height: 20px"></div>
    <h1>Рекомендуемый градус наклона солнечной панели относительно горизонта</h1>
    <div style="width: 100px; height: 20px"></div>
    <table>
      <tr>
        <Angle v-for="ind in index"
               v-bind:angle="angle_arr[ind]"
               v-bind:month="month[ind]"
        />
      </tr>
    </table>
    <div style="width: 100px; height: 50px"></div>
  <GlobAngle v-bind:angle_glob="angle_gl"/>
  </div>
</template>

<script>
import Graph from "@/components/Calc_forms/graph";
import Angle from "@/components/Calc_forms/angle-iteration"
import GlobAngle from "@/components/Calc_forms/global-angle"

export default {
  components: {
    Graph,
    Angle,
    GlobAngle,
  },
  data() {
    return {
      index: [
        0,
        1,
        2,
        3,
        4,
        5,
        6,
        7,
        8,
        9,
        10,
        11,
      ],
      angle_arr: this.response_data['angle'].splice(0, 12),
      angle_gl: this.response_data['angle'].splice(0, 1),
      month: [
        "January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December"
      ],
    }
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