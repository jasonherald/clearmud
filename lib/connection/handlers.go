package handlers

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net"
)

func RequestHandler(conn net.Conn, out chan string) {
	defer close(out)
	for {
		line, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			return
		}
		log.Print(string(line))
	}
}

func SendData(conn net.Conn, in <-chan string) {
	defer conn.Close()
	for {
		message := <-in
		if message != "" {
			log.Print(message)
			io.Copy(conn, bytes.NewBufferString(message))
		}
	}
}
