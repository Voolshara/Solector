package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.Use(cors.Default())
	initializeRoutes()
	gin.SetMode(gin.ReleaseMode)
	router.Run(":4325")
}
