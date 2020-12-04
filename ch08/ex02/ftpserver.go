package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/textproto"
	"os"
	"strings"
)

func main() {
	ftpListener, err := net.Listen("tcp", ":21")
	if err != nil {
		log.Fatal(err)
	}
	ftpSender, err := net.Listen("tcp", ":20")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ftpListener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, ftpSender)
	}
}

func handleConn(c net.Conn, ftpSender net.Listener) {
	var transferType string
	_, err := c.Write([]byte("220\n"))
	if err != nil {
		fmt.Printf("%v", err)
		c.Close()
	}
	fmt.Println(transferType)
	workingdir := os.Getenv("HOME")
	var sc net.Conn

	for {
		r := textproto.NewReader(bufio.NewReader(c))
		str, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				log.Printf("Disconnect")
				return
			}
			log.Fatal(err)
			return
		}
		log.Println(str)
		cmds := strings.Split(str, " ")
		switch cmds[0] {
		case "CWD":
			os.Chdir(workingdir)
			err := os.Chdir(cmds[1])
			if err != nil {
				c.Write([]byte("550\n"))
				break
			}
			workingdir, err = os.Getwd()
			if err != nil {
				c.Write([]byte("550\n"))
				break
			}
			c.Write([]byte("250\n"))
		case "HELP":
			log.Print("nyasu")
		case "LIST":
			c.Write([]byte("125\n"))
			log.Println("start LIST")
			os.Chdir(workingdir)
			if len(cmds) >= 2 {
				os.Chdir(cmds[1])
			}
			now, _ := os.Getwd()
			files, err := ioutil.ReadDir(now)
			log.Println(now)
			if err != nil {
				c.Write([]byte("550\n"))
			}
			for _, file := range files {
				t := "file"
				if file.IsDir() {
					t = "dir"
				}
				sc.Write([]byte(fmt.Sprintf("%15s  %10s %10d \n", file.Name(), t, file.Size())))
			}
			sc.Close()
			c.Write([]byte("226\n"))
		case "USER":
			log.Print(cmds)
			_, err = c.Write([]byte("331\n"))
		case "PASS":
			if len(cmds) != 2 {
				c.Write([]byte("530\n"))
			} else {
				if cmds[1] == "neko" {
					_, err = c.Write([]byte("230\n"))
					if err != nil {
						fmt.Errorf("%v", err)
						c.Write([]byte("530\n"))
					}
				}
			}
		case "EPSV":
			log.Println("hoge")
			c.Write([]byte("229 Entering Extended Passive Mode (|||20|) \n"))
			sc, err = ftpSender.Accept()
			if err != nil {
				c.Write([]byte("550 \n"))
				sc.Close()
			}
		case "PWD", "XPWD":
			c.Write([]byte(fmt.Sprintf("257 \"%s\" is the current directory\n", workingdir)))

		case "QUIT":
			c.Write([]byte("221 bye\n"))
		case "RETR":
			c.Write([]byte("125\n"))
			os.Chdir(workingdir)
			f, err := os.Open(cmds[1])
			if err != nil {
				log.Printf("FileOpen Error: %v", err)
				c.Write([]byte("550\n"))
			}
			io.Copy(sc, f)
			f.Close()
			sc.Close()
			c.Write([]byte("226\n"))

		case "STOR":
			c.Write([]byte("125\n"))
			os.Chdir(workingdir)
			f, err := os.Create(cmds[1])
			if err != nil {
				c.Write([]byte("550\n"))
			}
			f.Close()
			io.Copy(f, sc)
			sc.Close()
			c.Write([]byte("226\n"))
		case "FEAT", "PASV", "LPRT", "LPSV":
			c.Write([]byte("502\n"))
		case "SYST":
			c.Write([]byte("215 UNIX\n"))
		case "TYPE":
			switch cmds[1] {
			case "I":
				transferType = "I"
				c.Write([]byte("200\n"))
			case "A":
				transferType = "A"
				c.Write([]byte("200\n"))
			default:
				c.Write([]byte("502\n"))
			}
		case "SIZE":
			c.Write([]byte("202\n"))
		default:
			c.Write([]byte("202\n"))
		}
	}
}
