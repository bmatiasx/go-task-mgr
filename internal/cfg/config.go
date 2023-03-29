package cfg

import "os"

type Config struct {
	URL                string
	APIKey             string
	Token              string
	AppPort            string
	ToDoListId         string
	DoingListId        string
	BugLabelId         string
	MaintenanceLabelId string
	ResearchLabelId    string
	TestLabelId        string
}

func Setup() Config {
	conf := Config{
		URL:                os.Getenv("TRELLO_CARDS_URL"),
		APIKey:             os.Getenv("TRELLO_API_KEY"),
		Token:              os.Getenv("TRELLO_TOKEN"),
		AppPort:            os.Getenv("APP_PORT"),
		ToDoListId:         os.Getenv("TO_DO_LIST_ID"),
		DoingListId:        os.Getenv("DOING_LIST_ID"),
		BugLabelId:         os.Getenv("BUG_LABEL_ID"),
		MaintenanceLabelId: os.Getenv("MAINTENANCE_LABEL_ID"),
		ResearchLabelId:    os.Getenv("RESEARCH_LABEL_ID"),
		TestLabelId:        os.Getenv("TEST_LABEL_ID"),
	}
	return conf
}
