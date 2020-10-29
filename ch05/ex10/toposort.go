package main

import (
	"fmt"
)

var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":  {"discrete math": true},
	"databases":        {"data structures": true},
	"discrete math":    {"intro to programming": true},
	"formal languages": {"discrete math": true},
	"networks":         {"operating systems": true},
	"operating systems": {
		"data structures":       true,
		"computer organization": true,
	},
	"programming languages": {
		"data structures":       true,
		"computer organization": true,
	},
}

func main() {
	or := topoSort(prereqs)
	for i, course := range or {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	if isTopologicalOrdered(or) {
		fmt.Print("valid topology")
	} else {
		fmt.Print("invalid topology")
	}
}

func topoSort(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item, st := range items {
			if st {
				if !seen[item] {
					seen[item] = true
					visitAll(m[item])
					order = append(order, item)
				}
			}
		}
	}
	keys := map[string]bool{}
	for key := range m {
		keys[key] = true
	}
	visitAll(keys)
	return order
}

func isTopologicalOrdered(ts []string) bool {
	nodes := make(map[string]int)

	for i, course := range ts {
		nodes[course] = i
	}

	for course, i := range nodes {
		for prereq := range prereqs[course] {
			if i < nodes[prereq] {
				return false
			}
		}
	}
	return true
}
