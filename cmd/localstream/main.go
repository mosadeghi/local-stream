package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mosadeghi/local-stream/internal/admin"
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

	// Scan files in /movies
	videoDirs := []string{"D:\\Videos\\videos"}

	files, err := util.ListVideoFiles(videoDirs)
	if err != nil {
		log.Fatal("File scan failed:", err)
	}

	// Sync files with database
	if err := db.SyncMoviesWithDB(files); err != nil {
		log.Println("DB sync failed:", err)
	}
	router := gin.Default()

	router.Static("/static", "./web/static")

	router.LoadHTMLGlob("web/templates/*.html")

	adminGroup := router.Group("/admin", admin.BasicAuthMiddleware())
	adminGroup.GET("/", admin.ShowAdminPanel)
	adminGroup.POST("/update", admin.UpdateMovieMetadata)

	router.GET("/", func(c *gin.Context) {
		movies, err := db.GetAllMovies()
		if err != nil {
			log.Println("DB fetch failed:", err)
			movies = []db.Movie{}
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "LocalStream Home",
			"message": "Available Movies",
			"movies":  movies,
		})
	})

	router.GET("/movie/:id", public.ShowMoviePage)
	router.GET("/stream/:id", public.StreamVideo)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
