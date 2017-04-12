package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net"
)

func requestHandler(conn net.Conn, out chan string) {
	defer close(out)
	for {
		line, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			return
		}
		log.Print(string(line))
	}
}

func sendData(conn net.Conn, in <-chan string) {
	defer conn.Close()
	for {
		message := <-in
		if message != "" {
			log.Print(message)
			io.Copy(conn, bytes.NewBufferString(message))
		}
	}
}

func loadMotd() string {
	content, err := ioutil.ReadFile("MOTD")
	if err != nil {
		panic(1)
	}
	return string(content)
}

func main() {
	psock, err := net.Listen("tcp", ":5000")
	if err != nil {
		return
	}

	//var motd = loadMotd()

	for {
		conn, err := psock.Accept()
		if err != nil {
			return
		}

		channel := make(chan string)
		go requestHandler(conn, channel)
		go sendData(conn, channel)
		channel <- loadMotd()
	}
}
