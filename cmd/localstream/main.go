package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./web/static")

	router.LoadHTMLGlob("web/templates/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "LocalStream Home",
			"message": "Welcome to LocalStream!",
		})
	})

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
