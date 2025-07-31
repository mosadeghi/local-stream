package public

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mosadeghi/local-stream/internal/db"
	"github.com/mosadeghi/local-stream/internal/util"
)

const movieDir = "D:\\Videos\\videos"

func StreamVideo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.String(http.StatusBadRequest, "Invalid movie ID")
		return
	}

	// Fetch movie by ID
	movie, err := db.GetMovieByID(uint(id))
	if err != nil {
		c.String(http.StatusNotFound, "Movie not found")
		return
	}

	filePath := filepath.Join(movieDir, filepath.Clean(movie.FileName))

	file, err := os.Open(filePath)
	if err != nil {
		c.String(http.StatusNotFound, "File not found")
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not get file info")
		return
	}

	fileSize := stat.Size()
	rangeHeader := c.GetHeader("Range")

	if rangeHeader == "" {
		// No Range header: serve the whole file
		c.Header("Content-Type", "video/mp4")
		c.Header("Content-Length", fmt.Sprint(fileSize))
		c.Status(http.StatusOK)
		io.Copy(c.Writer, file)
		return
	}

	// Handle Range requests
	start, end := util.ParseRange(rangeHeader, fileSize)
	if start < 0 {
		c.Header("Content-Range", fmt.Sprintf("bytes */%d", fileSize))
		c.Status(http.StatusRequestedRangeNotSatisfiable)
		return
	}

	// Set headers
	contentLength := end - start + 1
	c.Header("Content-Type", "video/mp4")
	c.Header("Accept-Ranges", "bytes")
	c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
	c.Header("Content-Length", fmt.Sprint(contentLength))
	c.Status(http.StatusPartialContent)

	// Stream the requested byte range
	file.Seek(start, io.SeekStart)
	io.CopyN(c.Writer, file, contentLength)
}
