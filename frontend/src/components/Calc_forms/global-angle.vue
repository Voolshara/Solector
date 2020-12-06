<template>
  <div>
    <h1>Средний годовой градус наклона</h1>
    <table>
      <td>
        <tr style="width: 10px; height: 50px;"></tr>
        <tr>
          <span class="angle">{{ angle_glob[0] }}°</span>
        </tr>
      </td>
      <td>
        <tr>
          <img src="../../assets/solar2.png" alt="Sorry, image not supported" class="sol">
        </tr>
        <tr>
          <img src="../../assets/solar-panel2.png" alt="Sorry, image not supported" class="im">
        </tr>
      </td>
    </table>
  </div>
</template>

<script>
export default {
  props: {
    angle_glob: {
      type: Number,
      default: null,
    }
  },
  name: "global-angle"
}
</script>

<style>
.im {
  width: 75px;
  height: 75px;
}

.angle {
  color: #000000; /* Красный цвет символов */
  font-size: 300%; /* Размер шрифта в процентах */
}


.sol {
  width: 30px;
  height: 30px;
}
</style>