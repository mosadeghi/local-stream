package public

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mosadeghi/local-stream/internal/db"
)

func ShowMoviePage(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.String(http.StatusBadRequest, "Invalid movie ID")
		return
	}

	movie, err := db.GetMovieByID(uint(id))
	if err != nil {
		c.String(http.StatusNotFound, "Movie not found")
		return
	}

	c.HTML(http.StatusOK, "movie.html", gin.H{
		"movie": movie,
	})
}

func HomePage(c *gin.Context) {
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
}

func ListMovies(c *gin.Context) {
	movies, err := db.GetAllMovies()
	if err != nil {
		log.Println("DB fetch failed:", err)
		movies = []db.Movie{}
	}

	c.JSON(http.StatusOK, gin.H{
		"movies": movies,
	})
}

func MovieDetails(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid movie ID"})
		return
	}

	movie, err := db.GetMovieByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"movie": movie,
	})
}

type UpdateMovieRequest struct {
	Title    *string `json:"title"`
	Year     *int    `json:"year"`
	Director *string `json:"director"`
	Summary  *string `json:"summary"`
}

func UpdateMovieMetadata(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid movie ID"})
		return
	}

	movie, err := db.GetMovieByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
		return
	}

	// Update metadata
	var req UpdateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Title != nil {
		movie.Title = *req.Title
	}
	if req.Year != nil {
		movie.Year = *req.Year
	}
	if req.Director != nil {
		movie.Director = *req.Director
	}
	if req.Summary != nil {
		movie.Summary = *req.Summary
	}

	// Save to DB
	if err := db.DB.Save(movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update movie"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func UploadMoviePoster(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid movie ID"})
		return
	}

	movie, err := db.GetMovieByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Movie not found"})
		return
	}

	file, err := c.FormFile("poster")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Poster file is required"})
		return
	}

	// Save poster file
	dst := filepath.Join("web/static/posters", file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save poster"})
		return
	}

	// Update poster path in DB
	movie.PosterPath = "posters/" + file.Filename
	if err := db.DB.Save(movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update movie poster"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Poster uploaded successfully",
		"posterPath": movie.PosterPath,
	})
}
