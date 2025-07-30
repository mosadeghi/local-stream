package public

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const movieDir = "D:\\Videos\\videos"

func StreamVideo(c *gin.Context) {
	filename := c.Param("filename")
	filePath := filepath.Join(movieDir, filepath.Clean(filename))

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

	// Parse Range header (e.g., "bytes=1000-")
	const prefix = "bytes="
	if !strings.HasPrefix(rangeHeader, prefix) {
		c.Header("Content-Range", fmt.Sprintf("bytes */%d", fileSize))
		c.Status(http.StatusRequestedRangeNotSatisfiable)
		return
	}

	rangeSpec := strings.TrimPrefix(rangeHeader, prefix)
	parts := strings.Split(rangeSpec, "-")
	start, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil || start < 0 || start >= fileSize {
		c.Status(http.StatusRequestedRangeNotSatisfiable)
		return
	}

	end := fileSize - 1
	if len(parts) == 2 && parts[1] != "" {
		if parsedEnd, err := strconv.ParseInt(parts[1], 10, 64); err == nil && parsedEnd < fileSize {
			end = parsedEnd
		}
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
