package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	apiKey := os.Getenv("OMDB_KEY")
	clone := flag.NewFlagSet("clone", flag.ExitOnError)
	switch os.Args[1] {
	case "clone":
		endpoint := "https://omdbapi.com"
		u, err := url.Parse(endpoint)
		if err != nil {
			log.Printf("url parse Error:%v", err)
			return
		}
		str := clone.String("title", "", "target url")
		clone.Parse(os.Args[2:])

		q := u.Query()
		q.Set("apikey", apiKey)
		q.Set("t", *str)
		u.RawQuery = q.Encode()
		ret, err := cloneUrl(u.String())
		if err != nil {
			log.Printf("Clone Error: %v", err)
			return
		}
		err = savePicture(ret, os.Stdout)
		if err != nil {
			log.Printf("Get Poster Error: %v", err)
			return
		}
		fmt.Printf("Get Pic from %s", ret)
	}
}

func cloneUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("HTTP get error %v", err)
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body error %v", err)
		return "", err
	}
	var res interface{}
	if err := json.Unmarshal(bodyBytes, &res); err != nil {
		log.Printf("unmarshal error %v, %v", err, bodyBytes)
		return "", err
	}
	s := res.(map[string]interface{})["Poster"].(string)

	return s, nil
}

func savePicture(url string, w io.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("HTTP get error %v", err)
		return err
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
	return nil
}
