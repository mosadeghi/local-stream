package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mosadeghi/local-stream/internal/public"
	"github.com/mosadeghi/local-stream/internal/util"
)

const MoviesPath = "D:\\Videos\\videos"

func main() {
	router := gin.Default()

	router.Static("/static", "./web/static")

	router.LoadHTMLGlob("web/templates/*.html")

	router.GET("/", func(c *gin.Context) {
		movies, err := util.ListVideoFiles(MoviesPath)
		if err != nil {
			log.Println("Failed to scan movies:", err)
			movies = []string{}
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "LocalStream Home",
			"message": "Available Movies",
			"movies":  movies,
		})
	})

	router.GET("/stream/:filename", public.StreamVideo)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
