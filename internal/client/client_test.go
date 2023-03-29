package client

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/bmatiasx/go-task-mgr/internal/cfg"
	"github.com/bmatiasx/go-task-mgr/internal/model"
	"github.com/stretchr/testify/assert"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func TestClient_CreateIssue(t *testing.T) {

	issueJSON := `{
	"id": "6423991687731e2e9e1fec60",
	"idBoard": "63bdd2e8fdf46c026cf9aff2",
	"idList": "63bdd2e8fdf46c026cf9aff9",
	"url": "https://example.com/c/gpHVOuR7/65-check-api-key"
	}`

	reqString := "https://example.com/1/cards?idList=1&key=ABC123&token=123QWE"

	// Mock http client response
	httpClient := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		if req.URL.String() != reqString {
			t.Errorf("expected request to be %s, got %s", reqString, req.URL.String())
		}
		if req.Method != "POST" {
			t.Errorf("expected request method to be POST, got %s", req.Method)
		}
		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected request Content-Type header to be application/json, got %s", req.Header.Get("Content-Type"))
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(issueJSON)),
		}
	})}

	config := cfg.Config{
		URL:                "https://example.com",
		APIKey:             "ABC123",
		Token:              "123QWE",
		AppPort:            ":3000",
		ToDoListId:         "1",
		DoingListId:        "2",
		BugLabelId:         "10",
		MaintenanceLabelId: "11",
		ResearchLabelId:    "12",
		TestLabelId:        "13",
	}

	issue := model.Issue{
		Type:        "issue",
		Title:       "CD player not working",
		Description: "Change laser reader of CD player",
	}

	c := New(config)
	c.client = httpClient

	card, err := c.CreateIssue(issue)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.NotEmptyf(t, card.Id, "Id is empty")
	assert.NotEmptyf(t, card.Url, "URL is empty")
	assert.NotEmptyf(t, card.BoardId, "BoardId is empty")
	assert.NotEmptyf(t, card.ListId, "ListId is empty")
}

func TestClient_CreateBug(t *testing.T) {

	bugJSON := `{
	"id": "6423991687731e2e9e1fec60",
	"idBoard": "63bdd2e8fdf46c026cf9aff2",
	"idList": "63bdd2e8fdf46c026cf9aff9",
	"url": "https://example.com/c/gpHVOuR7/66-check-api-key"
	}`

	reqString := "https://example.com/1/cards?idList=2&key=ABC123&token=123QWE"

	// Mock http client response
	httpClient := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		if req.URL.String() != reqString {
			t.Errorf("expected request to be %s, got %s", reqString, req.URL.String())
		}
		if req.Method != "POST" {
			t.Errorf("expected request method to be POST, got %s", req.Method)
		}
		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected request Content-Type header to be application/json, got %s", req.Header.Get("Content-Type"))
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(bugJSON)),
		}
	})}

	config := cfg.Config{
		URL:                "https://example.com",
		APIKey:             "ABC123",
		Token:              "123QWE",
		AppPort:            ":3000",
		ToDoListId:         "1",
		DoingListId:        "2",
		BugLabelId:         "10",
		MaintenanceLabelId: "11",
		ResearchLabelId:    "12",
		TestLabelId:        "13",
	}

	bug := model.Bug{
		Type:        "Fuel indicator malfunction",
		Description: "Fuel level indicator not working properly",
	}

	c := New(config)
	c.client = httpClient

	card, err := c.CreateBug(bug)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.NotEmptyf(t, card.Id, "Id is empty")
	assert.NotEmptyf(t, card.Url, "URL is empty")
	assert.NotEmptyf(t, card.BoardId, "BoardId is empty")
	assert.NotEmptyf(t, card.ListId, "ListId is empty")
}

func TestClient_CreateTask(t *testing.T) {

	bugJSON := `{
	"id": "6423991687731e2e9e1fec60",
	"idBoard": "63bdd2e8fdf46c026cf9aff2",
	"idList": "63bdd2e8fdf46c026cf9aff9",
	"url": "https://example.com/c/gpHVOuR7/66-check-api-key"
	}`

	reqString := "https://example.com/1/cards?idList=1&key=ABC123&token=123QWE"

	// Mock http client response
	httpClient := &http.Client{Transport: RoundTripFunc(func(req *http.Request) *http.Response {
		if req.URL.String() != reqString {
			t.Errorf("expected request to be %s, got %s", reqString, req.URL.String())
		}
		if req.Method != "POST" {
			t.Errorf("expected request method to be POST, got %s", req.Method)
		}
		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected request Content-Type header to be application/json, got %s", req.Header.Get("Content-Type"))
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(bugJSON)),
		}
	})}

	config := cfg.Config{
		URL:                "https://example.com",
		APIKey:             "ABC123",
		Token:              "123QWE",
		AppPort:            ":3000",
		ToDoListId:         "1",
		DoingListId:        "2",
		BugLabelId:         "10",
		MaintenanceLabelId: "11",
		ResearchLabelId:    "12",
		TestLabelId:        "13",
	}

	task := model.Task{
		Type:     "Maintenance",
		Title:    "Keys cleaning",
		Category: "Program a keyboard cleaning every week",
	}

	c := New(config)
	c.client = httpClient

	card, err := c.CreateTask(task)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.NotEmptyf(t, card.Id, "Id is empty")
	assert.NotEmptyf(t, card.Url, "URL is empty")
	assert.NotEmptyf(t, card.BoardId, "BoardId is empty")
	assert.NotEmptyf(t, card.ListId, "ListId is empty")
}
