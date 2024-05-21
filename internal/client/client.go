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

	"github.com/bmatiasx/go-task-mgr/internal/cfg"
	"github.com/bmatiasx/go-task-mgr/internal/model"
)

const cardsPath = "1/cards"

//type Connector interface {
//	Do(req *http.Request) (*http.Response, error)
//}

type Client struct {
	URL     string
	APIKey  string
	Token   string
	AppPort string
	TaskIds
	LabelIds
	client *http.Client
}

type TaskIds struct {
	ToDoListId  string
	DoingListId string
	BugLabelId  string
}

type LabelIds struct {
	MaintenanceLabelId string
	ResearchLabelId    string
	TestLabelId        string
}

func New(cfg cfg.Config) *Client {
	c := Client{
		URL:     cfg.URL,
		APIKey:  cfg.APIKey,
		Token:   cfg.Token,
		AppPort: cfg.AppPort,
		TaskIds: TaskIds{
			ToDoListId:  cfg.ToDoListId,
			DoingListId: cfg.DoingListId,
			BugLabelId:  cfg.BugLabelId,
		},
		LabelIds: LabelIds{
			MaintenanceLabelId: cfg.MaintenanceLabelId,
			ResearchLabelId:    cfg.ResearchLabelId,
			TestLabelId:        cfg.TestLabelId,
		},
		client: &http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
	}
	return &c
}

func (c *Client) CreateIssue(request model.Issue) (*model.Card, error) {
	url := fmt.Sprintf("%s/%s?idList=%s&key=%s&token=%s", c.URL, cardsPath, c.ToDoListId, c.APIKey, c.Token)
	log.Printf("creating an issue with Trello API with url: %s", url)

	payload := map[string]string{
		"name": request.Title,
		"desc": request.Description,
	}
	issueResp := model.Card{}

	err := c.call(payload, &issueResp, http.MethodPost, url)
	if err != nil {
		log.Printf("error while creating an issue")
		return nil, fmt.Errorf("error: %s", err.Error())
	}
	return &issueResp, nil
}

func (c *Client) CreateBug(request model.Bug) (*model.Card, error) {
	url := fmt.Sprintf("%s/%s?idList=%s&key=%s&token=%s", c.URL, cardsPath, c.DoingListId, c.APIKey, c.Token)
	bugTitle := makeBugTitle()
	log.Printf("creating a bug with Trello API with url: %s \nand title: %s", url, bugTitle)

	payload := map[string]string{
		"name":     bugTitle,
		"desc":     request.Description,
		"idLabels": c.BugLabelId,
	}
	bugResp := model.Card{}

	err := c.call(payload, &bugResp, http.MethodPost, url)
	if err != nil {
		log.Printf("error while creating a bug")
		return nil, fmt.Errorf("error: %s", err.Error())
	}
	return &bugResp, nil
}

func (c *Client) CreateTask(request model.Task) (*model.Card, error) {
	url := fmt.Sprintf("%s/%s?idList=%s&key=%s&token=%s", c.URL, cardsPath, c.ToDoListId, c.APIKey, c.Token)
	log.Printf("creating a task with Trello API with url: %s", url)

	label := c.setLabel(request.Category)

	payload := map[string]string{
		"name":     request.Title,
		"desc":     fmt.Sprintf("Belongs to category %s", request.Category),
		"idLabels": label,
	}
	taskResp := model.Card{}

	err := c.call(payload, &taskResp, http.MethodPost, url)
	if err != nil {
		log.Printf("error while creating a task")
		return nil, fmt.Errorf("error: %s", err.Error())
	}
	return &taskResp, nil
}

func (c *Client) call(request interface{}, response *model.Card, httpMethod string, url string) error {

	b, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("error marshaling request")
	}

	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return fmt.Errorf("error creating request, %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request, %w", err)
	}

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body, %w", err)
	}

	switch {
	case !c.isSuccess(resp.StatusCode):
		log.Printf("response error, code: %v", resp.StatusCode)
		return fmt.Errorf("error returned from external API")
	default:
		log.Printf("successful client response. code: %v", resp.StatusCode)
		_ = json.Unmarshal(output, &response)
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

func (c *Client) setLabel(category string) string {
	categoryToLabel := map[string]string{
		"Maintenance": c.MaintenanceLabelId,
		"Research":    c.ResearchLabelId,
		"Test":        c.TestLabelId,
	}
	return categoryToLabel[category]
}
