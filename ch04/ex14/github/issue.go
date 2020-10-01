package github

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

const listIssuesURL = "https://api.github.com/repos/%s/%s/issues"

type IssuesListResult struct {
	Issues []*Issue
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

func GetIssues(owner string, path string) IssuesListResult {
	url := fmt.Sprintf(listIssuesURL, owner, path)
	resp, err := http.Get(url)
	if err != nil {
		return IssuesListResult{}
	}
	var ilr IssuesListResult
	if err := json.NewDecoder(resp.Body).Decode(&(ilr.Issues)); err != nil {
		return IssuesListResult{}
	}
	return ilr

}

func (ilr IssuesListResult) PrintAsHTMLTable(w http.ResponseWriter) {
	var ilt = template.Must(template.New("issuelist").Parse(`
	<h1>issues</h1>
	<table>
	<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
	</tr>
	{{range .Issues}}
	<tr>
	  <td><a href='issue/{{.Number}}'>{{.Number}}</a></td>
	  <td>{{.State}}</td>
	  <td>{{.User.Login}}</td>
	  <td><a href='issue/{{.Number}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))
	if err := ilt.Execute(w, ilr); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
}
