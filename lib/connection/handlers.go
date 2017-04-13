package handlers

import (
	combats "../combat"
	objects "../objects"
	"bufio"
	"bytes"
	"io"
	"net"
	"strings"
)

//RequestHandler Handles all incoming data
func RequestHandler(conn net.Conn, out chan string, user objects.User) {
	defer close(out) // close the connection when we are done
	for {
		line, err := bufio.NewReader(conn).ReadBytes('\n') // pull the line from the socket
		if err != nil {
			return
		}

		// look up the command which for now is just a movement
		// this code sucks... lol
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
		case "u":
			user.Move(4)
		case "d":
			user.Move(5)
		}
		out <- "Current Position: " + user.ToString() + "\n" // print this mess to the user for now
	}
}

//SendData handles all outgoing data
func SendData(conn net.Conn, in <-chan string) {
	defer conn.Close() // close when we are done
	for {
		message := <-in    // get the message we are writing from the channel
		if message != "" { // if it isn't blank
			io.Copy(conn, bytes.NewBufferString(message)) // write it to the socket
		}
	}
}
