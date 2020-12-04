package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	port := flag.Int("port", 8000, "clock server port")
	flag.Parse()
	tz := os.Getenv("TZ")
	if tz == "" {
		tz = "Asia/Tokyo"
	}
	loc, err := time.LoadLocation(tz)
	if err != nil {
		loc = time.FixedZone(tz, 9*60*60)
	}
	time.Local = loc
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05 MST\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
