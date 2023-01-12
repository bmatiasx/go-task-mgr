package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/bmatiasx/go-task-mgr/internal/model"
)

const (
	cardsPath   = "1/cards"
	toDoListId  = "63bdd2e8fdf46c026cf9aff9"
	doingListId = "63bdd2e8fdf46c026cf9affa"

	bugLabelId         = "63bdd2e87eabf59db1b0ad81"
	maintenanceLabelId = "63bdd2e87eabf59db1b0ad7b"
	researchLabelId    = "63bdd2e87eabf59db1b0ad87"
	testLabelId        = "63bdd2e87eabf59db1b0ad85"
)

type Connector interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	URL    string
	APIKey string
	Token  string
	Connector
}

func New(url, apiKey, token string) *Client {
	return &Client{
		URL:    url,
		APIKey: apiKey,
		Token:  token,
		Connector: &http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
	}
}

func (c *Client) CreateIssue(request model.Issue) (*model.Card, error) {
	url := fmt.Sprintf("%s/%s?idList=%s&key=%s&token=%s", c.URL, cardsPath, toDoListId, c.APIKey, c.Token)
	log.Printf("creating an issue with Trello API with url: %s", url)

	payload := map[string]string{
		"name": request.Title,
		"desc": request.Description,
	}
	issueResp := &model.Card{}

	err := c.Call(payload, issueResp, http.MethodPost, url)
	if err != nil {
		log.Printf("error while creating an issue")
		return nil, fmt.Errorf("error: %s", err.Error())
	}
	return issueResp, nil
}

func (c *Client) CreateBug(request model.Bug) (*model.Card, error) {
	url := fmt.Sprintf("%s/%s?idList=%s&key=%s&token=%s", c.URL, cardsPath, doingListId, c.APIKey, c.Token)
	bugTitle := makeBugTitle()
	log.Printf("creating a bug with Trello API with url: %s \nand title: %s", url, bugTitle)

	payload := map[string]string{
		"name":     bugTitle,
		"desc":     request.Description,
		"idLabels": bugLabelId,
	}
	bugResp := &model.Card{}

	err := c.Call(payload, bugResp, http.MethodPost, url)
	if err != nil {
		log.Printf("error while creating a bug")
		return nil, fmt.Errorf("error: %s", err.Error())
	}
	return bugResp, nil
}

func (c *Client) CreateTask(request model.Task) (*model.Card, error) {
	url := fmt.Sprintf("%s/%s?idList=%s&key=%s&token=%s", c.URL, cardsPath, toDoListId, c.APIKey, c.Token)
	log.Printf("creating a task with Trello API with url: %s", url)

	label := setLabel(request.Category)

	payload := map[string]string{
		"name":     request.Title,
		"desc":     fmt.Sprintf("Belongs to category %s", request.Category),
		"idLabels": label,
	}
	taskResp := &model.Card{}

	err := c.Call(payload, taskResp, http.MethodPost, url)
	if err != nil {
		log.Printf("error while creating a task")
		return nil, fmt.Errorf("error: %s", err.Error())
	}
	return taskResp, nil
}

func (c *Client) Call(request interface{}, response *model.Card, httpMethod string, url string) error {

	b, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("error marshaling request")
	}

	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.Connector.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request, %w", err)
	}

	output, err := ioutil.ReadAll(resp.Body)

	switch {
	case !c.isSuccess(resp.StatusCode):
		log.Printf("response error, code: %v", resp.StatusCode)
		return fmt.Errorf("error returned from external API")
	default:
		log.Printf("success. code: %v", resp.StatusCode)
		err = json.Unmarshal(output, &response)
	}

	return nil
}

func (c *Client) isSuccess(r int) bool {
	_, ok := c.successCodes()[r]
	return ok
}

func (c *Client) successCodes() map[int]string {
	return map[int]string{
		200: "success",
		202: "accepted",
		204: "no-content",
	}
}

func makeBugTitle() string {
	n := rand.Intn(999-0) + 0
	return fmt.Sprintf("bug-critical-%v", n)
}

func setLabel(category string) string {
	categoryToLabel := map[string]string{
		"Maintenance": maintenanceLabelId,
		"Research":    researchLabelId,
		"Test":        testLabelId,
	}
	return categoryToLabel[category]
}
