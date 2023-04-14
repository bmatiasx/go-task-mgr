package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

func TestHandleWelcome(t *testing.T) {
	mockTaskService := new(MockTaskService)
	handler := New(mockTaskService)

	// Given a get request to the welcome api
	req, err := http.NewRequest(http.MethodGet, "/api/v1/welcome", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	// When the handle method is called
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
