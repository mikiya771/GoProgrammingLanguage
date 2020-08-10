package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(getCommandName())
	str := joinArgsWithEnter(os.Args[1:])
	fmt.Println(str)
}

func getCommandName() string {
	return os.Args[0]
}

func getArgsWithIndex(args []string) []string {
	ret := []string{}
	for i, s := range args {
		ret = append(ret, strconv.Itoa(i+1)+" "+s)
	}
	return ret
}

func joinArgsWithEnter(args []string) string {
	str := ""
	strs := getArgsWithIndex(args)
	for _, s := range strs {
		str += s + "\n"
	}
	return str
}

func effectiveJoinArgsWithEnter(args []string) string {
	strs := getArgsWithIndex(args)
	str := strings.Join(strs, "\n")
	return str
}
