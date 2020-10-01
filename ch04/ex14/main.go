package main

import (
	"log"
	"mizuki/Develop/mikiya771/learn-language/GoProgrammingLangugage/ch04/ex14/github"
	"net/http"
	"strings"
)

type RepoPath struct {
	owner string
	repo  string
}

type RepoInfo struct {
	issues     github.IssuesListResult
	milestones github.MilestonesListResult
}

var repoInfoCaches = make(map[RepoPath]RepoInfo)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	paths := splitPath(r.URL.Path)
	rp := RepoPath{
		owner: paths[0],
		repo:  paths[1],
	}
	if _, ok := repoInfoCaches[rp]; !ok {
		repoInfoCaches[rp] = rp.GetGitHubAPI()
	}
	showPage(w, repoInfoCaches[rp])
}

func splitPath(path string) []string {
	paths := strings.Split(path[1:], "/")
	l := len(paths)
	if paths[l-1] == "" {
		return paths[:l-1]
	}
	return paths
}

func (rp RepoPath) GetGitHubAPI() RepoInfo {
	var ri RepoInfo
	ri.issues = github.GetIssues(rp.owner, rp.repo)
	ri.milestones = github.GetMilestones(rp.owner, rp.repo)
	return ri
}

func showPage(w http.ResponseWriter, ri RepoInfo) {
	ri.issues.PrintAsHTMLTable(w)
	ri.milestones.PrintAsHTMLTable(w)
}
