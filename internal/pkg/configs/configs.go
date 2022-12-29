package configs

import (
	"time"

	"main.go/internal/server/http"
)

// Configs struct handles all dependencies required for handling configurations
type Configs struct {
}

// HTTP returns the configuration required for HTTP
func (cfg *Configs) HTTP() (*http.Config, error) {
	return &http.Config{
		Port:         "8080",
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}, nil
}

// New returns an instance of Config with all the required dependencies initialized
func New() (*Configs, error) {
	return &Configs{}, nil
}
