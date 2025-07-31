package db

import (
	"fmt"
	"log"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dbPath string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Auto-migrate schema
	err = DB.AutoMigrate(&Movie{})
	if err != nil {
		return fmt.Errorf("auto migration failed: %w", err)
	}

	return nil
}

func SyncMoviesWithDB(paths []string) error {
	for _, path := range paths {
		var count int64
		if err := DB.Model(&Movie{}).Where("file_path = ?", path).Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			newMovie := Movie{
				FilePath: path,
				Title:    filepath.Base(path),
			}
			if err := DB.Create(&newMovie).Error; err != nil {
				log.Println("Failed to insert:", newMovie.Title, err)
			} else {
				log.Println("Inserted:", newMovie.Title)
			}
		}
	}
	return nil
}

func AddDummyMovie() error {
	dummy := Movie{
		FilePath:   "test_movie.mp4",
		Title:      "Test Movie",
		Year:       2024,
		Director:   "John Doe",
		Summary:    "A dummy movie for testing.",
		PosterPath: "posters/test_movie.jpg",
	}

	result := DB.Create(&dummy)
	return result.Error
}

func GetMovieByID(id uint) (*Movie, error) {
	var movie Movie
	result := DB.First(&movie, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &movie, nil
}

func GetAllMovies() ([]Movie, error) {
	var movies []Movie
	result := DB.Find(&movies)
	return movies, result.Error
}
