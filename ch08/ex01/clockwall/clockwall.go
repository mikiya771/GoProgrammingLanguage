package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
)

type time struct {
	index int
	time  string
}
type clockSetting struct {
	loc    string
	server string
}

func main() {
	flag.Parse()
	timer := []clockSetting{}
	for _, v := range flag.Args() {
		a := strings.Split(v, "=")
		if len(a) != 2 {
			fmt.Errorf("Parse Error %s is not valid", v)
			return
		}
		timer = append(timer, clockSetting{a[0], a[1]})
	}
	out := make(chan time, len(timer))
	times := make([]string, len(timer))
	for i, v := range timer {
		i, v := i, v
		go func() {
			conn, err := net.Dial("tcp", v.server)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			reader := bufio.NewReader(conn)
			for {
				bytes, _, err := reader.ReadLine()
				if err != nil {
					return
				}
				out <- time{i, fmt.Sprintf("%s: %s", v.loc, string(bytes))}
			}
		}()
	}
	for {
		t := <-out
		for i := 0; i < len(timer); i++ {
			times[t.index] = t.time
		}
		fmt.Print("\r")
		fmt.Print(strings.Join(times, " "))
	}
}
