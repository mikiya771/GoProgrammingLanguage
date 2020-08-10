package main

import (
	"strings"
	"testing"
)

func TestCommandName(t *testing.T) {
	actual := getCommandName()
	expect := "/ch01/ex01"
	if strings.Contains(expect, actual) {
		t.Errorf("unexpected output: %s", actual)
	}
}
