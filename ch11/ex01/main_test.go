package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	var tests = []struct {
		title  string
		input  string
		expect struct {
			counts  map[rune]int
			invalid int
		}
	}{
		{
			"ordinary case",
			"ABBCCCDDDDEEEEE",
			struct {
				counts  map[rune]int
				invalid int
			}{
				map[rune]int{'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5},
				0,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			actual_counts, actual_invalid, _, err := BufioToCharMap(bufio.NewReader(strings.NewReader(test.input)))
			if err != nil {
				t.Errorf("input: %s, unknown err: %v, expect: {count: %v, invalid: %d}", test.input, err, test.expect.counts, test.expect.invalid)
			}
			t.Run("check invalid number", func(t *testing.T) {
				if actual_invalid != test.expect.invalid {
					t.Errorf("Failed to Count Invalid Characters. Expect is %d, actual is %d", test.expect.invalid, actual_invalid)
				}
			})
			t.Run("check character maps", func(t *testing.T) {
				diffs := []string{}
				for key, val := range test.expect.counts {
					av, ok := actual_counts[key]
					if !ok {
						diffs = append(diffs, fmt.Sprintf("- [%s] expect: %d, but actual is not mapped", string(key), val))
					} else if val != av {
						diffs = append(diffs, fmt.Sprintf("~ [%s] expect: %d -> actual: %d", string(key), val, av))
					}
				}
				for key, val := range actual_counts {
					if _, ok := test.expect.counts[key]; !ok {
						diffs = append(diffs, fmt.Sprintf("+ [%s] not expected, but actual: %d", string(key), val))
					}
				}
				if len(diffs) != 0 {
					t.Errorf("character map has some differences,\n %s", strings.Join(diffs, "\n"))
				}
			})
		})
	}
}
