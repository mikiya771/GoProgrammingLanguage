package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stringTestCase struct {
	Name   string
	Input  string
	Prefix string
	Expect string
}

func TestMakeValidURL(t *testing.T) {
	testingTable := []stringTestCase{
		stringTestCase{"with valid url", "http://example.com", "http://", "http://example.com"},
		stringTestCase{"without http url", "example.com", "http://", "http://example.com"},
		stringTestCase{"without https url", "example.com", "https://", "https://example.com"},
	}
	for _, tt := range testingTable {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, tt.Expect, makeValidURL(tt.Input, tt.Prefix))
		})
	}
}

func TestGetResponseCode(t *testing.T) {
	OkResp := &http.Response{StatusCode: 200}
	NotFoundResp := &http.Response{StatusCode: 404}
	t.Run("200 OK", func(t *testing.T) { assert.Equal(t, getResponseCode(OkResp), 200) })
	t.Run("200 OK", func(t *testing.T) { assert.Equal(t, getResponseCode(NotFoundResp), 404) })
}
