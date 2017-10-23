package main

import (
	//handlers "./lib/connection"
	//objects "./lib/objects"
	sing "./lib/singletons"
	ticker "./lib/ticker"
	telnet "github.com/reiver/go-telnet"
	telsh "github.com/reiver/go-telnet/telsh"

	"io/ioutil"
	"io"
	"github.com/reiver/go-oi"
)

func main() {
	// new telnet nonsense
	shellHandler := telsh.NewShellHandler()

	// load in and parse the xml messages
	sing.GetMessagesInstance()

	content, err := ioutil.ReadFile("MOTD") // load the file contents for the MOTD
	if err != nil {                         // if reading the file wasn't successful
		panic(1) // oh dear god!!!!
	}

	shellHandler.WelcomeMessage = string(content)

	commandName     := "five"
	commandProducer := telsh.ProducerFunc(fiveProducer)

	shellHandler.Register(commandName, commandProducer)

	// spin up the ticker
	ticker.Start()
	defer ticker.Stop()

	addr := ":5000"
	if err := telnet.ListenAndServe(addr, shellHandler); nil != err {
		panic(err)
	}

	/*
	for { // loop forever


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
	*/
}


func fiveHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	oi.LongWriteString(stdout, "The number FIVE looks like this: 5\r\n")

	return nil
}

func fiveProducer(ctx telnet.Context, name string, args ...string) telsh.Handler {
	return telsh.PromoteHandlerFunc(fiveHandler)
}
