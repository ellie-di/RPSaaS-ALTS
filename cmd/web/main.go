package main

import (
	"log"
	"os"

	"main.go/internal/api"
	"main.go/internal/pkg/configs"
	"main.go/internal/server/http"
)

func main() {
	logger := log.New(os.Stderr, "alts:v0.0.1", 0)

	cfg, err := configs.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	httpCfg, err := cfg.HTTP()
	if err != nil {
		log.Fatal(err.Error())
	}

	api, err := api.New(*logger)
	if err != nil {
		log.Fatal(err.Error())
	}

	server, err := http.New(httpCfg, api)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	server.Start()
}
