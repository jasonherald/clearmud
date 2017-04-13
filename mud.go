package main

import (
	"io/ioutil"
	"net"

	handlers "./lib/connection"
	objects "./lib/objects"
	ticker "./lib/ticker"
)

func loadMotd() string {
	content, err := ioutil.ReadFile("MOTD") // load the file contents for the MOTD
	if err != nil {                         // if reading the file wasn't successulf
		panic(1) // oh dear god!!!!
	}
	return string(content) // return the string representation of the file's contents
}

func main() {
	psock, err := net.Listen("tcp", ":5000") // server socket declaration on port 5000
	if err != nil {                          // if the connection wasn't successful
		return // stop the program
	}

	// spin up the ticker
	ticker.Start()
	defer ticker.Stop()

	for { // loop forever
		conn, err := psock.Accept() // accept the incoming connection
		if err != nil {             // if the connection wasn't successful
			return // stop the program
		}

		//TODO: Need to retrieve the user's last position from storage
		user := objects.User{
			X: 0,
			Y: 0,
			Z: 0,
		}
		channel := make(chan string)                    // build a channel for concurrency
		go handlers.RequestHandler(conn, channel, user) // listen for incoming data
		go handlers.SendData(conn, channel)             // send data
		channel <- loadMotd()                           // send the motd
	}
}
