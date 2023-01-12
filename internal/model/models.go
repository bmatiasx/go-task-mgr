package model

type MasterTask struct {
	Type        string `json:"type,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
}

type Issue struct {
	Type        string
	Title       string
	Description string
}

type Bug struct {
	Type        string
	Description string
}

type Task struct {
	Type     string
	Title    string
	Category string
}

type Card struct {
	Id      string `json:"id"`
	Url     string `json:"url"`
	BoardId string `json:"idBoard"`
	ListId  string `json:"idList"`
}
