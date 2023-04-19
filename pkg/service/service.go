package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bmatiasx/go-task-mgr/internal/client"
	"github.com/bmatiasx/go-task-mgr/internal/model"
)

type Servicer interface {
	Welcome() string
	FilterTask(masterTask model.MasterTask) (map[string]string, error)
}

type strategy interface {
	createCard(task model.MasterTask) (map[string]string, error)
}

type TaskService struct {
	client.Client
	strategy
}

func New(client client.Client) *TaskService {
	return &TaskService{client, nil}
}

func (s *TaskService) Welcome() string {
	return "Welcome to the Card Service!"
}

func (s *TaskService) createCard(task model.MasterTask) (map[string]string, error) {

	return nil, nil
}

func (s *TaskService) FilterTask(masterTask model.MasterTask) (map[string]string, error) {

	err := validateRequest(masterTask)
	if err != nil {
		return nil, err
	}

	switch masterTask.Type {
	case "issue":
		issue := model.Issue{
			Type:        masterTask.Type,
			Title:       masterTask.Title,
			Description: masterTask.Description,
		}

		err := validateIssue(issue)
		if err != nil {
			return nil, err
		}

		// Call Trello API
		res, err := s.Client.CreateIssue(issue)
		if err != nil {
			return nil, err
		}
		log.Printf("Trello card created: [id: %s, url: %s, board_id: %s, list_id: %s]",
			res.Id, res.Url, res.BoardId, res.ListId)
		jsonResp := map[string]string{
			"message":     "card created",
			"type":        issue.Type,
			"title":       issue.Title,
			"description": issue.Description,
			"id":          res.Id,
			"url":         res.Url,
			"board_id":    res.BoardId,
			"list_id":     res.ListId,
		}
		return jsonResp, nil

	case "bug":
		bug := model.Bug{
			Type:        masterTask.Type,
			Description: masterTask.Description,
		}

		err := validateBug(bug)
		if err != nil {
			return nil, err
		}

		// Call Trello API
		res, err := s.Client.CreateBug(bug)
		if err != nil {
			return nil, err
		}
		log.Printf("Trello card created: [id: %s, url: %s, board_id: %s, list_id: %s]",
			res.Id, res.Url, res.BoardId, res.ListId)
		jsonResp := map[string]string{
			"message":     "card created",
			"type":        bug.Type,
			"description": bug.Description,
			"id":          res.Id,
			"url":         res.Url,
			"board_id":    res.BoardId,
			"list_id":     res.ListId,
		}
		return jsonResp, nil

	case "task":
		task := model.Task{
			Type:     masterTask.Type,
			Title:    masterTask.Title,
			Category: masterTask.Category,
		}

		err := validateTask(task)
		if err != nil {
			return nil, err
		}

		err = validateTaskCategory(masterTask.Category)
		if err != nil {
			return nil, err
		}
		// Call Trello API
		res, err := s.Client.CreateTask(task)
		if err != nil {
			return nil, err
		}
		log.Printf("Trello card created: [id: %s, url: %s, board_id: %s, list_id: %s]",
			res.Id, res.Url, res.BoardId, res.ListId)

		jsonResp := map[string]string{
			"message":  "card created",
			"type":     task.Type,
			"title":    task.Title,
			"category": task.Category,
			"id":       res.Id,
			"url":      res.Url,
			"board_id": res.BoardId,
			"list_id":  res.ListId,
		}
		return jsonResp, nil

	default:
		return map[string]string{"message": "non recognized task type"}, nil
	}
}

func validateRequest(task model.MasterTask) error {
	if len(task.Type) == 0 {
		log.Printf("error %+v. missing 'type' field", http.StatusBadRequest)
		return fmt.Errorf("empty 'type' field")
	}
	return nil
}

func validateIssue(issue model.Issue) error {
	if len(issue.Title) == 0 || len(issue.Description) == 0 {
		log.Printf("error %+v. empty 'title' or 'description' field", http.StatusBadRequest)
		return fmt.Errorf("error %+v: empty fields", http.StatusBadRequest)
	}
	return nil
}

func validateBug(issue model.Bug) error {
	if len(issue.Description) == 0 {
		log.Printf("error %+v. empty 'description' field", http.StatusBadRequest)
		return fmt.Errorf("error %+v: empty description", http.StatusBadRequest)
	}
	return nil
}

func validateTask(task model.Task) error {
	if len(task.Title) == 0 || len(task.Category) == 0 {
		log.Printf("error %+v. empty 'title' or 'category' field", http.StatusBadRequest)
		return fmt.Errorf("error %+v: empty fields", http.StatusBadRequest)
	}
	return nil
}

func validateTaskCategory(category string) error {
	c := []string{"Maintenance", "Research", "Test"}
	for _, v := range c {
		if v == category {
			return nil
		}
	}
	log.Printf("error %+v. invalid 'category' field", http.StatusBadRequest)
	return fmt.Errorf("error %+v: invalid category", http.StatusBadRequest)
}
