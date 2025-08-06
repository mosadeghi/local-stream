package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	AdminUsername string   `json:"admin_username"`
	AdminPassword string   `json:"admin_password"`
	MovieDirs     []string `json:"movie_dirs"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
