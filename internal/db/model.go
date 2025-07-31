package db

type Movie struct {
	ID         uint   `gorm:"primaryKey"`
	FileName   string `gorm:"uniqueIndex;not null"`
	Title      string
	Year       int
	Director   string
	Summary    string
	PosterPath string
}
