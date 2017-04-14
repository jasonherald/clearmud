package handlers

import (
	combats "../combat"
	objects "../objects"
	"bufio"
	"bytes"
	"io"
	"net"
	"strconv"
	"strings"
)

func RequestHandler(conn net.Conn, out chan string, user objects.User) {
	defer close(out)
	for {
		line, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			return
		}

		// look up the command which for now is just a movement
		str := string(line)
		str = strings.Trim(str, "\n")
		str = strings.Trim(str, "\r")
		switch str {
		case "n":
			user.Move(0)
		case "e":
			user.Move(1)
		case "w":
			user.Move(2)
		case "s":
			user.Move(3)
		case "fight":
			combats.Fight()
		}
		out <- "Current Position: " + user.ToString() + "\n"
		out <- strconv.FormatInt(combats.Fight(), 10) + "\n"

		//dot  your mom
	}
}

func SendData(conn net.Conn, in <-chan string) {
	defer conn.Close()
	for {
		message := <-in
		if message != "" {
			io.Copy(conn, bytes.NewBufferString(message))
		}
	}
}
