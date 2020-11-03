package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

/*
const ApiUrl = "https://power.larc.nasa.gov" +
	"/cgi-bin/v1/DataAccess.py" +
	"?&request=execute" +
	"&tempAverage=CLIMATOLOGY" +
	"&identifier=SinglePoint" +
	"&parameters=ALLSKY_SFC_SW_DWN" +
	"&userCommunity=SSE" +
	//"&lon=0&lat=0" +
	"&outputList=JSON" +
	"&user=DOCUMENTATION"
*/
func selection(nameCalc string, dataCalc map[string]string) (float64, float64) {
	if nameCalc == "dev-calc" {

		insolation, angle, err := NasaRequest(coords(dataCalc["coords"]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(insolation)
		fmt.Println(angle)

	} else if nameCalc == "power-calc" {

	} else {
		log.Fatal("Err Type Calculator")
	}
	return 3.0, 4.0
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
	fmt.Println(buf)
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
