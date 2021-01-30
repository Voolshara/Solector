<template>
  <div class="main">
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
          <td><input type="number" min="0" pattern="^[0-9]+$" v-model="form.kettle_kol"> Шт.</td>
          <td><input type="number" min="0" pattern="^[0-9]+$" v-model="form.kettle_wt"> Вт</td>
          <td><input type="number" min="0" pattern="^[0-9]+$" v-model="form.kettle_hour"> часов</td>
        </tr>
        <tr>
          <td>Холодильник</td>
          <td><input type="number" min="0" v-model="form.fridge_kol"> Шт.</td>
          <td><input type="number" min="0" v-model="form.fridge_wt"> Вт</td>
          <td><input type="number" min="0" v-model="form.fridge_hour"> часов</td>
        </tr>
        <tr>
          <td>Лампы</td>
          <td><input type="number" min="0" v-model="form.led_kol"> Шт.</td>
          <td><input type="number" min="0" v-model="form.led_wt"> Вт</td>
          <td><input type="number" min="0" v-model="form.led_hour"> часов</td>
        </tr>
        <tr>
          <td>Мультиварка</td>
          <td><input type="number" min="0" v-model="form.cooker_kol"> Шт.</td>
          <td><input type="number" min="0" v-model="form.cooker_wt"> Вт</td>
          <td><input type="number" min="0" v-model="form.cooker_hour"> часов</td>
        </tr>
        <tr>
          <td>Обогреватель</td>
          <td><input type="number" min="0" v-model="form.heater_kol"> Шт.</td>
          <td><input type="number" min="0" v-model="form.heater_wt"> Вт</td>
          <td><input type="number" min="0" v-model="form.heater_hour"> часов</td>
        </tr>
        <tr>
          <td>Стиральная машина</td>
          <td><input type="number" min="0" v-model="form.washing_kol"> Шт.</td>
          <td><input type="number" min="0" v-model="form.washing_wt"> Вт</td>
          <td><input type="number" min="0" v-model="form.washing_hour"> часов</td>
        </tr>
        <tr>
          <td>Телевизор</td>
          <td><input type="number" min="0" v-model="form.tv_kol"> Шт.</td>
          <td><input type="number" min="0" v-model="form.tv_wt"> Вт</td>
          <td><input type="number" min="0" v-model="form.tv_hour"> часов</td>
        </tr>
        <tr>
          <td>Ноутбук</td>
          <td><input type="number" min="0" v-model="form.laptop_kol"> Шт.</td>
          <td><input type="number" min="0" v-model="form.laptop_wt"> Вт</td>
          <td><input type="number" min="0" v-model="form.laptop_hour"> часов</td>
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
        <div style="width: 100%; height: 50px"></div>
        <div class="wrap">
          <button type="submit" class="button">Подобрать</button>
        </div>
      </table>
    </form>
    <div style="width: 1500px; height: 50px"></div>
    <div v-if="response_data['message']==='No Data'" style="width: 1500px; height: 600px">
      <h2>На данный момент у нас нет информации о инсоляции в этом регионе</h2>
    </div>
    <div v-else-if="this.response_data['power_arr'].length > 1">
      <response v-bind:response_data="response_data" class="response"/>
    </div>
    <div v-else-if="request_flag">
      <h2>Считаем...</h2>
      <loader class="loader"/>
    </div>
    <div v-else style="width: 1500px; height: 600px; margin-right: auto; margin-left: auto">
      <h2>Введите данные</h2>
    </div>
    <div style="width: 1500px; height: 80px;"></div>
  </div>

</template>

<script>
import form_head from '@/components/Calc_forms/form-head'
import {loadYmap} from 'vue-yandex-maps'
import Loader from "@/components/Calc_forms/loader";
import Response from "@/components/Calc_forms/form-response";


export default {
  components: {
    Loader,
    form_head,
    Response,
  },
  methods: {
    async jsonSubmit() {
      switch ('') {

        case this.form.cooker_hour:
          alert("Неправильное значение поле времени мультиварки")
              break
        case this.form.cooker_wt:
          alert("Неправильное значение поле мощности мультиварки")
          break
        case this.form.cooker_kol:
          alert("Неправильное значение поле колличества мультиварки")
          break

        case this.form.fridge_hour:
          alert("Неправильное значение поле времени холодильника")
          break
        case this.form.fridge_wt:
          alert("Неправильное значение поле мощности холодильника")
          break
        case this.form.fridge_kol:
          alert("Неправильное значение поле колличества холодильника")
          break

        case this.form.tv_hour:
          alert("Неправильное значение поле времени телевизора")
          break
        case this.form.tv_wt:
          alert("Неправильное значение поле мощности телевизора")
          break
        case this.form.tv_kol:
          alert("Неправильное значение поле колличества телевизора")
          break

        case this.form.kettle_hour:
          alert("Неправильное значение поле времени электрочайника")
          break
        case this.form.kettle_wt:
          alert("Неправильное значение поле мощности электрочайника")
          break
        case this.form.kettle_kol:
          alert("Неправильное значение поле колличества электрочайника")
          break

        case this.form.laptop_hour:
          alert("Неправильное значение поле времени ноутбука")
          break
        case this.form.laptop_wt:
          alert("Неправильное значение поле мощности ноутбука")
          break
        case this.form.laptop_kol:
          alert("Неправильное значение поле колличества ноутбука")
          break

        case this.form.led_hour:
          alert("Неправильное значение поле времени ламп")
          break
        case this.form.led_wt:
          alert("Неправильное значение поле мощности ламп")
          break
        case this.form.led_kol:
          alert("Неправильное значение поле колличества ламп")
          break

        case this.form.heater_hour:
          alert("Неправильное значение поле времени обогревателя")
          break
        case this.form.heater_wt:
          alert("Неправильное значение поле мощности обогревателя")
          break
        case this.form.heater_kol:
          alert("Неправильное значение поле колличества обогревателя")
          break

        case this.form.washing_hour:
          alert("Неправильное значение поле времени стиральной машины")
          break
        case this.form.washing_wt:
          alert("Неправильное значение поле мощности стиральной машины")
          break
        case this.form.washing_kol:
          alert("Неправильное значение поле колличества стиральной машины")
          break

        default:
          if (document.getElementById("mark_coord").value.length === 0) {
            alert("Установите метку на карте")
          } else {
            this.request_flag = 1
            this.response_data = {
              'power_arr': 0,
            }
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
                  this.response_data = data
                  console.log('Success:', data);
                })
                .catch((error) => {
                  this.response_data = error
                  console.error('Error:', error);
                });
          }
      }
    }
  },
  data() {
    return {
      request_flag: 0,
      response_data: {
        'power_arr': 0,
      },
      form: {
        kettle_kol: '1',
        kettle_wt: '360',
        kettle_hour: '1',
        fridge_kol: '1',
        fridge_wt: '150',
        fridge_hour: '6',
        led_kol: '5',
        led_wt: '8',
        led_hour: '12',
        cooker_kol: '0',
        cooker_wt: '500',
        cooker_hour: '2',
        heater_kol: '1',
        heater_wt: '1000',
        heater_hour: '5',
        washing_kol: '1',
        washing_wt: '500',
        washing_hour: '8',
        tv_kol: '1',
        tv_wt: '300',
        tv_hour: '5',
        laptop_kol: '1',
        laptop_wt: '60',
        laptop_hour: '10',
        coords: "",
      },
      "apiURL": "https://engine.solector.ru",
      //"apiURL": "http://192.168.1.82:4325/",
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

form {
  margin-left: auto;
  margin-right: auto;
}

.response {
  margin-right: auto;
  margin-left: auto;
}

.main {
  background-color: #ffffff;
}

.wrap {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.button {
  min-width: 300px;
  min-height: 60px;
  font-family: 'Nunito', sans-serif;
  font-size: 22px;
  text-transform: uppercase;
  letter-spacing: 1.3px;
  font-weight: 700;
  color: #313133;
  background: #4FD1C5;
  background: linear-gradient(90deg, rgba(129, 230, 217, 1) 0%, rgba(79, 209, 197, 1) 100%);
  border: none;
  border-radius: 1000px;
  box-shadow: 12px 12px 24px rgba(79, 209, 197, .64);
  transition: all 0.3s ease-in-out 0s;
  cursor: pointer;
  outline: none;
  position: relative;
  padding: 10px;
}

button::before {
  content: '';
  border-radius: 1000px;
  min-width: calc(300px + 12px);
  min-height: calc(60px + 12px);
  border: 6px solid #00FFCB;
  box-shadow: 0 0 60px rgba(0, 255, 203, .64);
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  opacity: 0;
  transition: all .3s ease-in-out 0s;
}

.button:hover, .button:focus {
  color: #313133;
  transform: translateY(-6px);
}

button:hover::before, button:focus::before {
  opacity: 1;
}

button::after {
  content: '';
  width: 30px;
  height: 30px;
  border-radius: 100%;
  border: 6px solid #00FFCB;
  position: absolute;
  z-index: -1;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  animation: ring 1.5s infinite;
}

button:hover::after, button:focus::after {
  animation: none;
  display: none;
}

@keyframes ring {
  0% {
    width: 30px;
    height: 30px;
    opacity: 1;
  }
  100% {
    width: 300px;
    height: 300px;
    opacity: 0;
  }
}

.loader {
  height: 400px;
  margin-top: 200px;
  width: 500px;
  margin-left: auto;
  margin-right: auto;
}

td {
  padding: 10px 10px;
}

</style>