package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	githubAPI "mizuki/Develop/mikiya771/learn-language/GoProgrammingLangugage/ch04/ex11/githubAPI"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/labstack/gommon/log"
)

const PREFIX_TYPE = "ISSUE"

var defaultSetting []string = []string{"repo:mikiya771/GoProgrammingLanguage"}

func main() {
	// listIssue := flag.NewFlagSet("list", flag.ExitOnError)
	getIssue := flag.NewFlagSet("get", flag.ExitOnError)
	createIssue := flag.NewFlagSet("create", flag.ExitOnError)
	deleteIssue := flag.NewFlagSet("delete", flag.ExitOnError)
	updateIssue := flag.NewFlagSet("update", flag.ExitOnError)

	var user githubAPI.Credentials
	user.Query()
	switch os.Args[1] {
	case "list":
		ret, err := githubAPI.SearchIssues(defaultSetting)
		if err != nil {
		}
		for _, item := range ret.Items {
			fmt.Printf("#%-5d %9.9s %55s\n", item.Number, item.User.Login, item.Title)
		}
	case "get":
		num := getIssue.Int("number", 1, "issue number")
		getIssue.Parse(os.Args[2:])
		ret, err := githubAPI.SearchIssues(defaultSetting)
		if err != nil {
		}
		for _, item := range ret.Items {
			if item.Number == *num {
				fmt.Printf("number:#%d \n", item.Number)
				fmt.Printf("Title: %s \n", item.Title)
				fmt.Printf("State: %s \n", item.State)
				fmt.Printf("CreatedAt: %s \n", item.CreatedAt)
				fmt.Println("body: ")
				fmt.Printf("%s\n", item.Body)
			}
		}
		return
	case "create":
		str := createIssue.String("title", "", "Issue Title")
		createIssue.Parse(os.Args[2:])
		b := openIssueFile()
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", "mikiya771", "GoProgrammingLanguage")
		body, err := json.Marshal(githubAPI.CreateIssue{Title: *str, Body: b})
		req, err := http.NewRequest("POST", url, bytes.NewReader(body))
		req.SetBasicAuth(user.Username, user.Password)
		req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
		if err != nil {
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
		}
		defer resp.Body.Close()
		var patchedIssue githubAPI.Issue
		if err := json.NewDecoder(resp.Body).Decode(&patchedIssue); err != nil {
			return
		}

		if patchedIssue.State == "close" {
			fmt.Println("Issue Closed")
		}
	case "delete":
		num := deleteIssue.Int("number", 0, "issue number")
		deleteIssue.Parse(os.Args[2:])
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d", "mikiya771", "GoProgrammingLanguage", *num)
		fmt.Println(url)
		body, err := json.Marshal(githubAPI.CloseIssue{State: "close"})
		req, err := http.NewRequest("PATCH", url, bytes.NewReader(body))
		req.SetBasicAuth(user.Username, user.Password)
		req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
		if err != nil {
			fmt.Printf("1: %v", err)
			return
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("2: %v", err)
			return
		}
		defer resp.Body.Close()
		var r io.Reader = resp.Body
		r = io.TeeReader(r, os.Stderr)
		var patchedIssue githubAPI.Issue
		fmt.Println(*num)
		if err := json.NewDecoder(r).Decode(&patchedIssue); err != nil {
			fmt.Printf("3: %v", err)
			return
		}

		if patchedIssue.State == "close" {
			fmt.Println("Issue Closed")
		}
	case "update":
		str := updateIssue.String("title", "", "Issue Title")
		num := updateIssue.Int("num", 0, "Issue number")
		updateIssue.Parse(os.Args[2:])
		b := openIssueFile()
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d", "mikiya771", "GoProgrammingLanguage", *num)
		fmt.Println(url)
		body, err := json.Marshal(githubAPI.UpdateIssue{Title: *str, Body: b, State: "Open"})
		req, err := http.NewRequest("PATCH", url, bytes.NewReader(body))
		req.SetBasicAuth(user.Username, user.Password)
		req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
		if err != nil {
			fmt.Printf("1: %v", err)
			return
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("2: %v", err)
			return
		}
		defer resp.Body.Close()
		var patchedIssue githubAPI.Issue
		if err := json.NewDecoder(resp.Body).Decode(&patchedIssue); err != nil {
			fmt.Printf("3: %v", err)
			return
		}

		if patchedIssue.State == "close" {
			fmt.Println("Issue Closed")
		}
	}
}

func openIssueFile() string {
	fp := filepath.Join("/tmp", fmt.Sprintf("%s_EDITMSG", PREFIX_TYPE))
	var err error
	err = initTempFile(fp)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	err = launchVim(fp)
	if err != nil {
		os.Exit(1)
	}
	bytes, err := ioutil.ReadFile(fp)
	if err != nil {
		panic(fmt.Errorf("Cannot read a temp file: %v", err))
	}
	return string(bytes)
}

func launchVim(args ...string) error {
	c := exec.Command("nvim", args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

func initTempFile(fp string) error {
	_, err := os.Stat(fp)
	if err != nil {
		if !os.IsNotExist(err) {
			err := os.Remove(fp)
			if err != nil {
				return err
			}
		}
	}
	f, err := os.Create(fp)
	if err != nil {
		fmt.Println("ho")
		return err
	}
	defer f.Close()

	return nil
}
