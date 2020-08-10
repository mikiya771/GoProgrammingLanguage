package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(getCommandName())
    for _, s := range getArgsWithIndex(os.Args[1:]){
        fmt.Println(s)
    }
}

func getCommandName() string {
	return os.Args[0]
}
func getArgsWithIndex(args []string)[]string{
    ret := []string{}
    for i, s := range args{
        ret = append(ret, strconv.Itoa(i + 1)+ " " + s)
    }
    return ret
}
