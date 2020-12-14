package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type NasaResponse struct {
	Features []struct {
		Properties struct {
			Parameter struct {
				Insolation []float64 `json:"SI_EF_OPTIMAL"`
				Angle      []float64 `json:"SI_EF_OPTIMAL_ANG"`
			}
		}
	}
}

const ApiUrl = "https://power.larc.nasa.gov" +
	"/cgi-bin/v1/DataAccess.py" +
	"?&request=execute" +
	"&tempAverage=CLIMATOLOGY" +
	"&identifier=SinglePoint" +
	"&parameters=SI_EF_TILTED_SURFACE" +
	"&userCommunity=SSE" +
	//"&lon=0&lat=0" +
	"&outputList=JSON" +
	"&user=DOCUMENTATION"

func selection(nameCalc string, dataCalc map[string]string) (string, []float64, []float64, int64, int64,
																[]float64, string, string, string, string) {
	if nameCalc == "dev-calc" {
		insolation, angle, err := NasaRequest(coords(dataCalc["coords"]))
		if len(insolation) == 0 {
			var arrErr []float64
			return "None", arrErr, arrErr, 0, 0, arrErr, "", "", "", "NoData"
		} else if len(angle) == 0 {
			var arrErr []float64
			return "None", arrErr, arrErr, 0, 0, arrErr, "", "", "", "NoData"
		} else {
			if err != nil {
				log.Fatal(err)
			}
			panelName, powerArr, cost, kol, abs, img_link, company_name, panel_link := search(count(dataCalc), insolation)
			absArr := []float64{abs, abs, abs, abs, abs, abs, abs, abs, abs, abs, abs, abs, abs}
			return panelName, powerArr, angle, cost, kol, absArr, img_link, company_name, panel_link, "None"
		}
	} else if nameCalc == "power-calc" {
		var arrErr []float64
		return "None", arrErr, arrErr, 0, 0, arrErr, "", "", "", "NoData"
	} else {
		log.Fatal("Err Type Calculator")
	}
	var arrErr []float64
	return "None", arrErr, arrErr, 0, 0, arrErr, "", "", "", "NoData"
}

func search(abs float64, ins []float64) (string, []float64, int64, int64, float64, string, string, string) {
	min := minInsolation(ins[0 : len(ins)-1])
	db, err := sql.Open("sqlite3", "./back_solector.db")
	if err != nil {
		fmt.Println("sql.Open err", err)
	}
	rows, err := db.Query("SELECT * FROM panels")
	if err != nil {
		fmt.Println("db.Query err", err)
	}
	var db_id int64
	var db_name string
	var db_cost int64
	var db_power float64
	var db_efficiency float64
	var img_link string
	var company string
	var panel_link string
	//--------------------------------------------------------------------------------------
	var minArr [2]int64
	var minName string
	var minCoef float64
	var minImg string
	var coef float64
	var min_comp string
	var min_panel string
	i := 0
	for rows.Next() {
		err = rows.Scan(&db_id, &db_name, &db_cost, &db_power, &db_efficiency, &img_link, &company, &panel_link)
		coef = (db_power / 1000) * (db_efficiency / 100)
		powerPanel := min * coef
		kol := cleverInt(abs / powerPanel)
		cost := kol * db_cost
		if err != nil {
			fmt.Println("rows.Scan err", err)
		}
		if i == 0 {
			minArr[0] = cost // minArr[0] - cost
			minArr[1] = kol
			minCoef = coef
			minName = db_name
			minImg = img_link
			min_comp = company
			min_panel = panel_link
		} else {
			if cost < minArr[0] {
				minArr[0] = cost // minArr[0] - cost
				minArr[1] = kol
				minCoef = coef
				minName = db_name
				minImg = img_link
				min_comp = company
				min_panel = panel_link
			}
		}
		i++
	}
	var OutputArr []float64
	for i := 0; i < len(ins); i++ {
		OutputArr = append(OutputArr, ins[i]*minCoef*float64(minArr[1]))
	}
	rows.Close() //good habit to close
	db.Close()
	return minName, OutputArr, minArr[0], minArr[1], abs, minImg, min_comp, min_panel
}

func NasaRequest(lat float64, lon float64) ([]float64, []float64, error) {
	url := fmt.Sprintf("%s&lat=%f&lon=%f", ApiUrl, lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	var NasaResp NasaResponse
	err = json.Unmarshal(buf.Bytes(), &NasaResp)
	if err != nil {
		return nil, nil, err
	}
	return NasaResp.Features[0].Properties.Parameter.Insolation, NasaResp.Features[0].Properties.Parameter.Angle, nil
}

func coords(st string) (float64, float64) {
	a := strings.Split(st, ",")
	b, err := strconv.ParseFloat(a[0][:4], 64)
	if err != nil {
		log.Fatal("err cords", err)
	}
	c, err := strconv.ParseFloat(a[1][:4], 64)
	if err != nil {
		log.Fatal("err cords", err)
	}
	return b, c
}

func cleverInt(x float64) int64 {
	ost := x - float64(int64(x))
	if ost >= 0.00 {
		return int64(x) + 1
	} else {
		return int64(x)
	}

}

func minInsolation(arr []float64) float64 {
	min := arr[0]
	for _, value := range arr {
		if min > value {
			min = value
		}
	}
	return min
}

func cleverFloat(s string, bitSize int) (float64, error) {
	if s == ""{
		return 0.00, errors.New("err parsing ''")
	} else {
		va, err := strconv.ParseFloat(s, bitSize)
		if err != nil{
			return 0.00, err
		}
		return va, err
	}
}

func count(data map[string]string) float64 {
	cookerKol, err := cleverFloat(data["cooker_kol"], 64)
	if err != nil {
		log.Fatal("cookerKol Err", err)
	}
	cookerWt, err := cleverFloat(data["cooker_wt"], 64)
	if err != nil {
		log.Fatal("cookerWt Err", err)
	}
	cookerHour, err := cleverFloat(data["cooker_hour"], 64)
	if err != nil {
		log.Fatal("cookerHour Err", err)
	}
	fridgeKol, err := cleverFloat(data["fridge_kol"], 64)
	if err != nil {
		log.Fatal("fridgeKol Err", err)
	}
	fridgeWt, err := cleverFloat(data["fridge_wt"], 64)
	if err != nil {
		log.Fatal("fridgeWt Err", err)
	}
	fridgeHour, err := cleverFloat(data["fridge_hour"], 64)
	if err != nil {
		log.Fatal("fridgeHour Err", err)
	}
	heaterKol, err := cleverFloat(data["heater_kol"], 64)
	if err != nil {
		log.Fatal("heaterKol Err", err)
	}
	heaterWt, err := cleverFloat(data["heater_wt"], 64)
	if err != nil {
		log.Fatal("heaterWt Err", err)
	}
	heaterHour, err := cleverFloat(data["heater_hour"], 64)
	if err != nil {
		log.Fatal("heaterHour Err", err)
	}
	kettleKol, err := cleverFloat(data["kettle_kol"], 64)
	if err != nil {
		log.Fatal("kettleKol Err", err)
	}
	kettleWt, err := cleverFloat(data["kettle_wt"], 64)
	if err != nil {
		log.Fatal("kettleWt Err", err)
	}
	kettleHour, err := cleverFloat(data["kettle_hour"], 64)
	if err != nil {
		log.Fatal("kettleHour Err", err)
	}
	laptopKol, err := cleverFloat(data["laptop_kol"], 64)
	if err != nil {
		log.Fatal("laptopKol Err", err)
	}
	laptopWt, err := cleverFloat(data["laptop_wt"], 64)
	if err != nil {
		log.Fatal("laptopWt Err", err)
	}
	laptopHour, err := cleverFloat(data["laptop_hour"], 64)
	if err != nil {
		log.Fatal("laptopHour Err", err)
	}
	ledKol, err := cleverFloat(data["led_kol"], 64)
	if err != nil {
		log.Fatal("ledKol Err", err)
	}
	ledWt, err := cleverFloat(data["led_wt"], 64)
	if err != nil {
		log.Fatal("ledWt Err", err)
	}
	ledHour, err := cleverFloat(data["led_hour"], 64)
	if err != nil {
		log.Fatal("ledHour Err", err)
	}
	tvKol, err := cleverFloat(data["tv_kol"], 64)
	if err != nil {
		log.Fatal("tvKol Err", err)
	}
	tvWt, err := cleverFloat(data["tv_wt"], 64)
	if err != nil {
		log.Fatal("tvWt Err", err)
	}
	tvHour, err := cleverFloat(data["tv_hour"], 64)
	if err != nil {
		log.Fatal("tvHour Err", err)
	}
	washingKol, err := cleverFloat(data["washing_kol"], 64)
	if err != nil {
		log.Fatal("washingKol Err", err)
	}
	washingWt, err := cleverFloat(data["washing_wt"], 64)
	if err != nil {
		log.Fatal("washingWt Err", err)
	}
	washingHour, err := cleverFloat(data["washing_hour"], 64)
	if err != nil {
		log.Fatal("washingHour Err", err)
	}
	return ((cookerWt * cookerKol * (cookerHour / 24)) +
		(fridgeKol * fridgeWt * (fridgeHour / 24)) +
		(heaterKol * heaterWt * (heaterHour / 24)) +
		(kettleKol * kettleWt * (kettleHour / 24)) +
		(laptopKol * laptopWt * (laptopHour / 24)) +
		(ledKol * ledWt * (ledHour / 24)) +
		(tvKol * tvWt * (tvHour / 24)) +
		(washingKol * washingWt * (washingHour / 24))) / 1000
}
