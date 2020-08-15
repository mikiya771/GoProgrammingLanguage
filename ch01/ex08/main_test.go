package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
