package util

import (
	"io/fs"
	"path/filepath"
	"strings"
)

var videoExtensions = map[string]bool{
	".mp4":  true,
	".avi":  true,
	".mkv":  true,
	".mov":  true,
	".webm": true,
}

func ListVideoFiles(dirPath string) ([]string, error) {
	var videos []string

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // skip on error
		}
		if !d.IsDir() {
			ext := strings.ToLower(filepath.Ext(d.Name()))
			if videoExtensions[ext] {
				videos = append(videos, d.Name())
			}
		}
		return nil
	})

	return videos, err
}
