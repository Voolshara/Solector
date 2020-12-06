<template>
  <td>
    <tr>
      {{ month }}
    </tr>
    <tr>
      <table>
        <td style="width: 15px; height: 1px">
          <div ></div>
        </td>
        <td>
          <img src="../../assets/solar2.png" v-bind:alt="angle" class="solar">
        </td>
      </table>
    </tr>
    <tr>
      <img src="../../assets/solar-panel2.png" v-bind:alt="angle" class="img">
    </tr>
    <tr>
      <div style="width: 50px; height: 10px"></div>
    </tr>
    <tr >
      {{ angle }}°
    </tr>
  </td>
</template>

<script>
export default {
  props: {
    angle: {
      type: Number,
      default: null,
    },
    month: {
      type: String,
      default: null,
    },
  },
  name: "angle-iteration"
}
</script>

<style>
.img {
  width: 70px;
  height: 70px;
}

.solar {
  width: 30px;
  height: 30px;
}
</style>