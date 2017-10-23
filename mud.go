package main

import (
	"net"

	handlers "./lib/connection"
	objects "./lib/objects"
	sing "./lib/singletons"
	ticker "./lib/ticker"
)
import (
	telnet "github.com/reiver/go-telnet"
)

func main() {
	var handler telnet.Handler = telnet.EchoHandler

	err := telnet.ListenAndServe(":5555", handler)
	if nil != err {
		//@TODO: Handle this error better.
		panic(err)
	}

	psock, err := net.Listen("tcp", ":5000") // server socket declaration on port 5000
	if err != nil {                          // if the connection wasn't successful
		return // stop the program
	}

	// load in and parse the xml messages
	sing.GetMessagesInstance()

	// spin up the ticker
	ticker.Start()
	defer ticker.Stop()

	for { // loop forever
		conn, err := psock.Accept() // accept the incoming connection
		if err != nil {             // if the connection wasn't successful
			return // stop the program
		}

		channel := make(chan string) // build a channel for concurrency

		//TODO: Need to retrieve the user's last position from storage
		user := &objects.User{
			X:              0,
			Y:              0,
			Z:              0,
			C:              channel,
			Name:           "Jason",
			SentTimeUpdate: 0,
		}

		go handlers.RequestHandler(conn, channel, user) // listen for incoming data
		go handlers.SendData(conn, channel)             // send data
		sing.GetUsersInstance().Add(user)               // register this user in the singleton
	}
}
