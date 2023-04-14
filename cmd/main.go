package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bmatiasx/go-task-mgr/internal/cfg"
	"github.com/bmatiasx/go-task-mgr/internal/client"
	"github.com/bmatiasx/go-task-mgr/internal/controller"
	"github.com/bmatiasx/go-task-mgr/pkg/service"
)

func main() {
	fmt.Println("Welcome to Task Manager")

	log.SetFlags(0)
	config := cfg.Setup()

	clientConnector := *client.New(config)
	srv := service.New(clientConnector)

	mux := http.NewServeMux()
	mux.Handle("/", controller.New(srv))
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatalf("Service will be shutdown because an error occured: %+v", err.Error())
	}
}
