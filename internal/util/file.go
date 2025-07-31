package util

import (
	"os"
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

func ListVideoFiles(rootDirs []string) ([]string, error) {
	var videos []string
	for _, root := range rootDirs {
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}
			ext := strings.ToLower(filepath.Ext(path))
			if videoExtensions[ext] {
				normalized := filepath.ToSlash(path)
				videos = append(videos, normalized)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return videos, nil
}
