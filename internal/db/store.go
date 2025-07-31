package db

import (
	"fmt"

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

func AddDummyMovie() error {
	dummy := Movie{
		FileName:   "test_movie.mp4",
		Title:      "Test Movie",
		Year:       2024,
		Director:   "John Doe",
		Summary:    "A dummy movie for testing.",
		PosterPath: "posters/test_movie.jpg",
	}

	result := DB.Create(&dummy)
	return result.Error
}

func GetAllMovies() ([]Movie, error) {
	var movies []Movie
	result := DB.Find(&movies)
	return movies, result.Error
}
