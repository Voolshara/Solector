package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type js_struct struct {
	Name     string `json:"name"`
	DataForm map[string]string `json:"data"`
}



func initializeRoutes() {

	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"INdevices.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	})
	router.POST("/calcout", func(c *gin.Context) {
		var js_form js_struct
		err := c.BindJSON(&js_form)
		if err != nil{
			log.Fatal(err)
		}
		typeCalculate := js_form.Name
		dataCalculate := js_form.DataForm
		x, y := selection(typeCalculate, dataCalculate)
		fmt.Println(x)
		fmt.Print(y)
		c.JSON(http.StatusOK, gin.H {
			"message":"It's working",
		})
	})
	router.GET("/calcout", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message":"WHAT you doing?",
		})
	})
}

