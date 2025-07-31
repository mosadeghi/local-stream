package admin

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mosadeghi/local-stream/internal/db"
)

func ShowAdminPanel(c *gin.Context) {
	movies, _ := db.GetAllMovies()

	c.HTML(http.StatusOK, "admin.html", gin.H{
		"movies": movies,
	})
}

func UpdateMovieMetadata(c *gin.Context) {
	idStr := c.PostForm("id")
	title := c.PostForm("title")
	yearStr := c.PostForm("year")
	director := c.PostForm("director")
	summary := c.PostForm("summary")

	id, _ := strconv.Atoi(idStr)
	year, _ := strconv.Atoi(yearStr)

	movie, err := db.GetMovieByID(uint(id))
	if err != nil {
		c.String(http.StatusNotFound, "Movie not found")
		return
	}

	// Update metadata
	movie.Title = title
	movie.Year = year
	movie.Director = director
	movie.Summary = summary

	// Handle poster upload
	file, err := c.FormFile("poster")
	if err == nil {
		dst := filepath.Join("web/static/posters", file.Filename)
		if err := c.SaveUploadedFile(file, dst); err == nil {
			movie.PosterPath = "posters/" + file.Filename
		}
	}

	// Save to DB
	db.DB.Save(movie)
	c.Redirect(http.StatusSeeOther, "/admin")
}
