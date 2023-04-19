package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bmatiasx/go-task-mgr/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) Welcome() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockTaskService) FilterTask(_ model.MasterTask) (map[string]string, error) {
	args := m.Called()
	return args.Get(0).(map[string]string), args.Error(1)
}

func TestTaskHandler_TestHandleWelcome(t *testing.T) {
	mockTaskService := new(MockTaskService)
	handler := New(mockTaskService)

	// Given a get request to the welcome api
	req, err := http.NewRequest(http.MethodGet, "/api/v1/welcome", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	// When the HandleWelcome method is called
	mockTaskService.On("Welcome").Once().Return("Welcome message")
	handler.HandleWelcome(recorder, req)

	// Then check the status code and the message content
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	var res map[string]string
	err = json.Unmarshal(recorder.Body.Bytes(), &res)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %s", err)
	}

	expected := "Welcome message"
	if res["message"] != expected {
		t.Errorf("Expected message '%s', but got '%s'", expected, res["message"])
	}
}

func TestTaskHandler_HandleTask(t *testing.T) {
	mockTaskService := new(MockTaskService)
	handler := New(mockTaskService)
	inputReq := "{\n    \"type\": \"task\",\n    \"title\": \"Brush keys on all boards\",\n    \"category\": \"Maintenance\"\n}"

	// Given an incoming task comes in the form of request
	req, err := http.NewRequest(http.MethodPost, "/", strings.NewReader(inputReq))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	// When the HandleTask method is invoked
	serviceRes := map[string]string{
		"board_id": "890uio",
		"category": "Maintenance",
		"id":       "123qwe",
		"list_id":  "asd456",
		"message":  "card created",
		"title":    "Brush keys on all panels",
		"type":     "task",
		"url":      "https://example.com/c/ueMYnIVX/69-brush-keys-on-all-boards",
	}

	mockTaskService.On("FilterTask").Return(serviceRes, nil)

	handler.HandleTask(recorder, req)

	// Then check that the expected object is valid
	var res map[string]string
	err = json.Unmarshal(recorder.Body.Bytes(), &res)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %s", err)
	}

	expectedRes := map[string]string{
		"board_id": "890uio",
		"category": "Maintenance",
		"id":       "123qwe",
		"list_id":  "asd456",
		"message":  "card created",
		"title":    "Brush keys on all panels",
		"type":     "task",
		"url":      "https://example.com/c/ueMYnIVX/69-brush-keys-on-all-boards",
	}

	if res["message"] != expectedRes["message"] {
		t.Errorf("Expected message '%s', but got '%s'", expectedRes["message"], res["message"])
	}
	if res["title"] != expectedRes["title"] {
		t.Errorf("Expected title '%s', but got '%s'", expectedRes["title"], res["title"])
	}
}
