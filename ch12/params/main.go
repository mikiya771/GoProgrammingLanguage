package main

import (
	"net/http"

	"gopl.io/ch12/params"
)

func search(resp http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels      []string `http:"1"`
		MaxResulsts int      `http:"max"`
		Exact       bool     `http:"x"`
	}
	data.MaxResulsts = 10
	if err := params.Unpack(req, &data); err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
}
