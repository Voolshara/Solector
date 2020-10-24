<template>
  <div>
    <br>
    <br>
    <form @submit.prevent="jsonSubmit">
      <table>
        <tr>
          <form_head
              v-for="head in form_head"
              v-bind:head="head"
          />
        </tr>
        <br>
        <tr>
          <td>Электрочайник</td>
          <td><input type="number" v-model="form.kettle_kol" value="0"> Шт.</td>
          <td><input type="number" v-model="form.kettle_wt" value="360"> Вт</td>
          <td><input type="number" v-model="form.kettle_hour" val> часов</td>
        </tr>
        <tr>
          <td>Холодильник</td>
          <td><input type="number" v-model="form.fridge_kol"> Шт.</td>
          <td><input type="number" v-model="form.fridge_wt"> Вт</td>
          <td><input type="number" v-model="form.fridge_hour"> часов</td>
        </tr>
        <tr>
          <td>Лампы</td>
          <td><input type="number" v-model="form.led_kol"> Шт.</td>
          <td><input type="number" v-model="form.led_wt"> Вт</td>
          <td><input type="number" v-model="form.led_hour"> часов</td>
        </tr>
        <tr>
          <td>Мультиварка</td>
          <td><input type="number" v-model="form.cooker_kol"> Шт.</td>
          <td><input type="number" v-model="form.cooker_wt"> Вт</td>
          <td><input type="number" v-model="form.cooker_hour"> часов</td>
        </tr>
        <tr>
          <td>Обогреватель</td>
          <td><input type="number" v-model="form.heater_kol"> Шт.</td>
          <td><input type="number" v-model="form.heater_wt"> Вт</td>
          <td><input type="number" v-model="form.heater_hour"> часов</td>
        </tr>
        <tr>
          <td>Стиральная машина</td>
          <td><input type="number" v-model="form.washing_kol"> Шт.</td>
          <td><input type="number" v-model="form.washing_wt"> Вт</td>
          <td><input type="number" v-model="form.washing_hour"> часов</td>
        </tr>
        <tr>
          <td>Телевизор</td>
          <td><input type="number" v-model="form.tv_kol"> Шт.</td>
          <td><input type="number" v-model="form.tv_wt"> Вт</td>
          <td><input type="number" v-model="form.tv_hour"> часов</td>
        </tr>
        <tr>
          <td>Ноутбук</td>
          <td><input type="number" v-model="form.laptop_kol"> Шт.</td>
          <td><input type="number" v-model="form.laptop_wt"> Вт</td>
          <td><input type="number" v-model="form.laptop_hour"> часов</td>
        </tr>
      </table>
      <br>
      <table>
        <tr>
          <td>
            <div id="map" style="width: 1500px; height: 500px"></div>
            <input id="mark_coord" type="text" v-model="form.coords" style="display: none;">
          </td>
        </tr>
        <input type="submit" value="Submit" style="height: 100px; width: 300px"/>
      </table>
      <br>
      <br>
      <br>
      <br>
      <br>
      <br>
    </form>
  </div>

</template>

<script>
import form_head from '@/components/Calc_forms/form-head'
import {loadYmap} from 'vue-yandex-maps'
import axios from 'axios';
import VeeValidate from 'vee-validate'

export default {
  components: {
    form_head,
  },
  methods: {
    async jsonSubmit() {
      this.form.coords = document.getElementById("mark_coord").value
      let js_form = {
        name: "dev-calc",
        data: this.form
      }
      await fetch(this.apiURL, {
        method: 'POST', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, *cors, same-origin
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, *same-origin, omit
        headers: {
          'Content-Type': 'application/json'
          // 'Content-Type': 'application/x-www-form-urlencoded',
        },
        redirect: 'follow', // manual, *follow, error
        referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
        body: JSON.stringify(js_form) // body data type must match "Content-Type" header
      })
          .then(response => response.json())
          .then(data => {
            console.log('Success:', data);
          })
          .catch((error) => {
            console.error('Error:', error);
          });
    }
  },
  data() {
    return {
      form: {
        kettle_kol: '0',
        kettle_wt: '360',
        kettle_hour: '1',
        fridge_kol: '0',
        fridge_wt: '150',
        fridge_hour: '6',
        led_kol: '0',
        led_wt: '8',
        led_hour: '12',
        cooker_kol: '0',
        cooker_wt: '500',
        cooker_hour: '2',
        heater_kol: '0',
        heater_wt: '1000',
        heater_hour: '5',
        washing_kol: '0',
        washing_wt: '500',
        washing_hour: '8',
        tv_kol: '0',
        tv_wt: '300',
        tv_hour: '5',
        laptop_kol: '0',
        laptop_wt: '60',
        laptop_hour: '10',
        coords: "",
      },
      "apiURL": "http://localhost:4325/calcout",
      "form_head": [
        "Электроприбор",
        "Количество",
        "Мощность",
        "Время работы за день",
      ],
    }
  },
  //------------------------------------------------------------------------------MAP--------------------------------
  async mounted() {
    const map_settings = {
      apiKey: "49ba2530-8257-44cb-b1cb-a19fb39e046c",
      lang: 'ru_RU',
      coordorder: 'latlong',
      version: '2.1'
    };
    await loadYmap({map_settings, debug: true});
    ymaps.ready(init);

    function init() {
      var myPlacemark,
          myMap = new ymaps.Map('map', {
            center: [55.753994, 37.622093],
            zoom: 7
          }, {
            searchControlProvider: 'yandex#search'
          });

      // Слушаем клик на карте.
      myMap.events.add('click', function (e) {
        var coords = e.get('coords');
        document.getElementById('mark_coord').value = coords;
        // Если метка уже создана – просто передвигаем ее.
        if (myPlacemark) {
          myPlacemark.geometry.setCoordinates(coords);
        }
        // Если нет – создаем.
        else {
          myPlacemark = createPlacemark(coords);
          myMap.geoObjects.add(myPlacemark);
          // Слушаем событие окончания перетаскивания на метке.
          myPlacemark.events.add('dragend', function () {
            getAddress(myPlacemark.geometry.getCoordinates());
          });
        }
        getAddress(coords);
      });

      // Создание метки.
      function createPlacemark(coords) {
        return new ymaps.Placemark(coords, {
          iconCaption: 'здесь солнечная электростаниця'
        }, {
          preset: 'islands#violetDotIconWithCaption',
          draggable: true
        });
      }

      // Определяем адрес по координатам (обратное геокодирование).
      function getAddress(coords) {
        myPlacemark.properties.set('iconCaption', 'солнечная электростанция');
        /*  ymaps.geocode(coords).then(function (res) {
            myPlacemark.properties
                .set({
                  "sfdgfhgjk":"sfdgfhg"
                });
          });*/
      }
    }
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
}

td {
  padding: 10px 10px;
}

</style>