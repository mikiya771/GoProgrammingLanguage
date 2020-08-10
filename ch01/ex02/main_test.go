package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stringsEquationTestcase struct {
	Input  []string
	Expect []string
}

func TestCommandName(t *testing.T) {
	testingTable := []stringsEquationTestcase{
		stringsEquationTestcase{Input: []string{}, Expect: []string{}},
		stringsEquationTestcase{Input: []string{"a"}, Expect: []string{"1 a"}},
		stringsEquationTestcase{Input: []string{"a", "b"}, Expect: []string{"1 a", "2 b"}},
	}
	actualName := getCommandName()
	expectName := "/ex02"
	assert.Contains(t, actualName, expectName, fmt.Sprintf("command name should countain %s", expectName))
	for _, tc := range testingTable {
		actual := getArgsWithIndex(tc.Input)
		assert.Exactly(t, tc.Expect, actual, "the two slices should be exactly euqal")
	}
}
