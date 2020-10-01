package github

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Milestone struct {
	State        string
	Title        string
	Description  string
	OpenIssues   int `json:"open_issues"`
	ClosedIssues int `json:"closed_issues"`
}

const milestonesURL = "https://api.github.com/repos/%s/%s/milestones"

type MilestonesListResult struct {
	Milestones []Milestone
}

func GetMilestones(owner string, repo string) MilestonesListResult {
	url := fmt.Sprintf(milestonesURL, owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return MilestonesListResult{}
	}
	var mlr MilestonesListResult
	if err := json.NewDecoder(resp.Body).Decode(&(mlr.Milestones)); err != nil {
		return MilestonesListResult{}
	}
	return mlr
}

func (mlr MilestonesListResult) PrintAsHTMLTable(w http.ResponseWriter) {
	var mlt = template.Must(template.New("milestoneList").Parse(`
	<h1>milestones</h1>
	<table>
	<tr style='text-align: left'>
	<th>Title</th>
	<th>State</th>
	</tr>
	{{range .Milestones}}
	<tr>
	  <td>{{.Title}}</td>
	  <td>{{.State}}</td>
	</tr>
	{{end}}
	</table>
	`))

	if err := mlt.Execute(w, mlr); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	}
}
