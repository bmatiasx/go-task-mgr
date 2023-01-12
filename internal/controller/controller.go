package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bmatiasx/go-task-mgr/internal/model"
	"github.com/bmatiasx/go-task-mgr/pkg/service"
)

const (
	welcome = "/api/v1/welcome"
	task    = "/"
)

type TaskHandler struct {
	service service.TaskService
}

func New(s service.TaskService) *TaskHandler {
	handler := TaskHandler{service: s}
	return &handler
}

func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == welcome:
		h.handleWelcome(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path == task:
		h.handleTask(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *TaskHandler) handleWelcome(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s API was called", r.URL)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(h.service.Welcome()))
}

func (h *TaskHandler) handleTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling new task")

	masterTask := unmarshalMasterTask(w, r)

	res, err := h.service.FilterTask(masterTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s", err.Error())))
		return
	}
	jsonRes, err := json.Marshal(res)

	if err != nil {
		log.Fatalf("error marshaling json response. %s", err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonRes)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s endpoint not found", r.URL)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func unmarshalMasterTask(w http.ResponseWriter, r *http.Request) model.MasterTask {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Error while reading request body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Unmarshal
	var masterTask model.MasterTask
	err = json.Unmarshal(b, &masterTask)
	if err != nil {
		log.Println("Error while unmarshalling request")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return masterTask
}
