package api

import (
	"log"
	"time"
)

var (
	now = time.Now()
)

// API holds all the dependencies required to expose APIs. And each API is a function with *API as its receiver
type API struct {
	logger log.Logger
}

// New returns a new instance of API with all the dependencies initialized
func New(l log.Logger) (*API, error) {
	return &API{
		logger: l,
	}, nil
}
