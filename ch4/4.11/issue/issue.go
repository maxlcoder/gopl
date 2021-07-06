package issue

import (
	"encoding/json"
	"net/http"
)

type Params struct {
	Owner  string
	Repo   string
	Number string
	Token  string
	Issue
}

type Issue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

const baseUrl = "https://api.github.com/repos/"

func (p Params) GetIssues() ([]Issue, error) {
	u := baseUrl + p.Owner + "/" + p.Repo + "/issues"
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func (p Params) GetIssue()  {
	
}
