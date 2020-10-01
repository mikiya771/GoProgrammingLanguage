package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	now := time.Now()
	fmt.Println("作られてから一ヶ月未満のissue")
	for _, item := range result.Items {
		if item.CreatedAt.AddDate(0, 1, 0).After(now) {
			fmt.Printf("#%-5d %9.9s %55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Println("作られてから一年未満のissue")
	for _, item := range result.Items {
		if item.CreatedAt.AddDate(0, 1, 0).Before(now) && item.CreatedAt.AddDate(1, 0, 0).After(now) {
			fmt.Printf("#%-5d %9.9s %55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Println("作られてから一年以上のissue")
	for _, item := range result.Items {
		if item.CreatedAt.AddDate(1, 0, 0).After(now) {
			fmt.Printf("#%-5d %9.9s %55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}
