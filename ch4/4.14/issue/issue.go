package issue

import (
	"encoding/json"
	"net/http"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []Issue
}

type Issue struct {
	Title string
	User User
	Milestone Milestone
}

type User struct {
	Login string
}

type Milestone struct {
	Title string
	Description string
}

func SearchIssue(params string) (*IssueSearchResult, error) {
	resp, err := http.Get(IssuesURL + "?q=" + params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result IssueSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}