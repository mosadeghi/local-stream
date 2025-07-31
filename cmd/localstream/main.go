package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mosadeghi/local-stream/internal/db"
	"github.com/mosadeghi/local-stream/internal/public"
	"github.com/mosadeghi/local-stream/internal/util"
)

const MoviesPath = "D:\\Videos\\videos"

func main() {
	// Init DB
	err := db.InitDatabase("metadata.db")
	if err != nil {
		panic("DB init failed: " + err.Error())
	}

	// Test DB:
	// if err := db.AddDummyMovie(); err != nil {
	// 	log.Println("Add dummy failed:", err)
	// }

	// if movies, err := db.GetAllMovies(); err != nil {
	// 	log.Println("Query failed:", err)
	// } else {
	// 	log.Printf("Found %d movie(s)\n", len(movies))
	// }

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

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
