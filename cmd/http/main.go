package main

import (
	"flag"
	"log"
	"net/http"
	"repartners/cmd/http/config"
	"repartners/cmd/http/handlers"
	"repartners/internal/pack/repository"
)

func main() {
	appConfig := initConfig()
	log.Printf("app initialized with config: %+v", *appConfig)

	packHandler := handlers.PackHandler{
		Repo: repository.NewMemoryRepo(),
	}

	http.HandleFunc("/pack_sizes", packHandler.Calculate)

	log.Fatal(http.ListenAndServe(appConfig.Port, nil))
}

// initializing the configuration of the app, reading all the flags / env settings.
func initConfig() *config.Config {
	port := flag.String("port", ":3000", "app port")
	flag.Parse()

	appConfig := config.Config{
		Port: *port,
	}

	return &appConfig
}
