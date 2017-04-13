package singletons

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"sync"
)

type messagesSingleton struct {
	MessageList []message `xml:"Message"`
}

type message struct {
	OnPlanet int
	When     int
	Body     string
}

var messagesInstance *messagesSingleton // single singleton instance
var mOnce sync.Once                     // threadsafe runner

//GetMessagesInstance returns the singleton instance
func GetMessagesInstance() *messagesSingleton {
	mOnce.Do(func() { // run once
		messagesInstance = &messagesSingleton{} // instantiate the singleton

		xmlFile, err := ioutil.ReadFile("./messages/time.xml")

		if err != nil {
			log.Println("Error opening file:", err)
			//return
		}

		xml.Unmarshal(xmlFile, messagesInstance)
	})
	return messagesInstance // return the instance we created (if we created one)
}
