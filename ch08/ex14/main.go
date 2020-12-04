package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	ch  chan<- string
	who string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli, _ := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			cli.ch <- "members:"
			for c, _ := range clients {
				cli.ch <- fmt.Sprintf("%s, ", c.who)
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	var who string
	input := bufio.NewScanner(conn)
	if input.Scan() {
		who = input.Text()
	}
	ch <- "You are" + who
	messages <- who + " has arrived"
	entering <- client{ch, who}

	timer := time.AfterFunc(time.Minute*5, func() {
		conn.Close()
	})
	for input.Scan() {
		timer.Stop()
		messages <- who + ": " + input.Text()
		timer = time.AfterFunc(time.Minute*5, func() {
			conn.Close()
		})
	}
	timer.Stop()
	leaving <- client{ch, who}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
