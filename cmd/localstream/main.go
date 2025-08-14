package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mosadeghi/local-stream/internal/admin"
	"github.com/mosadeghi/local-stream/internal/config"
	"github.com/mosadeghi/local-stream/internal/db"
	"github.com/mosadeghi/local-stream/internal/public"
	"github.com/mosadeghi/local-stream/internal/util"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Init DB
	err = db.InitDatabase("metadata.db")
	if err != nil {
		panic("DB init failed: " + err.Error())
	}

	// Scan files in /movies
	files, err := util.ListVideoFiles(cfg.MovieDirs)
	if err != nil {
		log.Fatal("File scan failed:", err)
	}

	// Sync files with database
	if err := db.SyncMoviesWithDB(files); err != nil {
		log.Println("DB sync failed:", err)
	}

	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*.html")
	router.Static("/static", "./web/static")

	router.GET("/", public.HomePage)
	router.GET("/stream/:id", public.StreamVideo)
	adminGroup := router.Group("/admin", admin.BasicAuthMiddleware(cfg))
	{
		adminGroup.GET("/", admin.ShowAdminPanel)
		adminGroup.POST("/update", admin.UpdateMovieMetadata)
	}

	api := router.Group("/api")
	{
		api_v1 := api.Group("/v1")
		{
			api_v1.GET("/movies", public.ListMovies)

			api_v1.GET("/movies/:id", public.MovieDetails)
			api_v1.PUT("/movies/:id", public.UpdateMovieMetadata)
			api_v1.POST("/movies/:id/poster", public.UploadMoviePoster)
		}
	}

	router.Static("/app", "./frontend/dist")

	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/app") {
			c.File("./frontend/dist/index.html")
			return
		}
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Oooooooops! Donno what you looking for!",
		})
	})

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
