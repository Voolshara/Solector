package main

import (
"fmt"
"github.com/gin-gonic/gin"
"log"
"net/http"
"strconv"
"strings"
)

type js_struct struct {
	Name     string `json:"name"`
	DataForm map[string]string `json:"data"`
}

func coord(x string) (int64, int64) {
	a := strings.Split(x, ",")
	b := strings.Split(a[0], ".")
	c := strings.Split(a[1], ".")
	return conversion(b[0]), conversion(c[0])
}
func conversion(x string) int64 {
	out, err := strconv.ParseInt(x[0:], 10, 64)
	if err != nil {
		fmt.Println("Err_Conversion  ", err)
	}
	if out < 0 {
		out = 0
	}
	return out
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
		fmt.Println(js_form)

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

