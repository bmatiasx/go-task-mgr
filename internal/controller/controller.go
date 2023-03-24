package controller

import (
	"encoding/json"
	"fmt"
	"io"
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
	_, err := w.Write([]byte(h.service.Welcome()))
	if err != nil {
		log.Fatalf("error writing json response. %s", err)
	}
}

func (h *TaskHandler) handleTask(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling new task")

	masterTask, err := unmarshalMasterTask(w, r)
	if err != nil {
		res := map[string]string{
			"error":   fmt.Sprintf("%+v", http.StatusBadRequest),
			"message": err.Error(),
		}
		jsonRes, err := json.Marshal(res)
		if err != nil {
			log.Fatalf("error marshaling json response. %s", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(jsonRes)
		if err != nil {
			log.Fatalf("error writing json response. %s", err)
		}
		return
	}

	res, err := h.service.FilterTask(masterTask)
	if err != nil {
		log.Fatalf("error creating task. %s", err)
		return
	}

	jsonRes, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("error marshaling json response. %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonRes)
	if err != nil {
		log.Fatalf("error writing json response. %s", err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s endpoint not found", r.URL)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("not found"))
	if err != nil {
		log.Fatalf("error writing json response. %s", err)
	}
}

func unmarshalMasterTask(w http.ResponseWriter, r *http.Request) (model.MasterTask, error) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error cloding body")
			return
		}
	}(r.Body)
	if err != nil {
		log.Println("Error while reading request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Unmarshal
	var masterTask model.MasterTask
	err = json.Unmarshal(b, &masterTask)
	if err != nil {
		log.Println("Error while unmarshalling request")
		return model.MasterTask{}, fmt.Errorf("error while unmarshalling request")
	}
	return masterTask, nil
}
