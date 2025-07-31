package public

import (
	"net/http"
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
