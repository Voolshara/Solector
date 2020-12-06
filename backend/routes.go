package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type jsStruct struct {
	Name     string            `json:"name"`
	DataForm map[string]string `json:"data"`
}

func initializeRoutes() {
	router.POST("/", func(c *gin.Context) {
		var js_form jsStruct
		err := c.BindJSON(&js_form)
		if err != nil {
			log.Fatal(err)
		}
		typeCalculate := js_form.Name
		dataCalculate := js_form.DataForm
		help_arr := []float64{1.00, 3.00}
		name, power, angle, cost, kol, abs, img_link, company_name, panel_link, data := selection(typeCalculate,
			dataCalculate)
		if data == "NoData" {
			c.JSON(http.StatusOK, gin.H{
				"message":   "No Data",
				"power_arr": help_arr,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"name":       name,
				"cost":       cost,
				"kol":        kol,
				"company":    company_name,
				"panel_link": panel_link,
				"img_link":   img_link,
				"power":      abs,
				"power_arr":  power,
				"angle":      angle,
				"message":    "All good",
			})
		}
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "WHAT you doing?",
		})
	})
}
