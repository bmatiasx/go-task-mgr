package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bmatiasx/go-task-mgr/internal/client"
	"github.com/bmatiasx/go-task-mgr/internal/controller"
	"github.com/bmatiasx/go-task-mgr/pkg/service"
)

// Ideally this should be in a config file
const (
	appPort        = ":3000"
	trelloCardsUrl = "https://api.trello.com"
	trelloAPIKey   = "0833ee5275c65f2e89c1cb983e54dd1a"
	trelloToken    = "ATTA12a0dcc2f794779e6830ad8e3306cb7db7d0ce685c18393931119ef5e37e098363A7FBE7"
)

type Config struct {
	URL    string
	APIKey string
	Token  string
}

func main() {
	fmt.Println("Welcome to Task Manager")

	cfg := Config{
		URL:    trelloCardsUrl,
		APIKey: trelloAPIKey,
		Token:  trelloToken,
	}
	clientConnector := *client.New(cfg.URL, cfg.APIKey, cfg.Token)
	srv := *service.New(clientConnector)

	mux := http.NewServeMux()
	mux.Handle("/", controller.New(srv))
	if err := http.ListenAndServe(appPort, mux); err != nil {
		log.Fatalf("Service will be shutdown because an error occured: %+v", err.Error())
	}
}
