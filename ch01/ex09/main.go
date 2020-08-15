package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const urlPrefix = "http://"

func main() {
	for _, url := range os.Args[1:] {
		url := makeValidURL(url, urlPrefix)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		_, err = io.Copy(os.Stdin, resp.Body)
		fmt.Println(getResponseCode(resp))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func makeValidURL(originalURL string, validPrefix string) string {
	ind := strings.Index(originalURL, validPrefix)
	if ind == 0 {
		return originalURL
	}
	return validPrefix + originalURL
}

func getResponseCode(resp *http.Response) int {
	return resp.StatusCode
}
