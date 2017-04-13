package messages

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

type TimeMessage struct {
	MessageList []Message `xml:"Message"`
}

type Message struct {
	OnPlanet int
	When     int
	Body     string
}

func Parse() {
	xmlFile, err := ioutil.ReadFile("./messages/time.xml")

	if err != nil {
		log.Println("Error opening file:", err)
		//return
	}

	var q TimeMessage
	xml.Unmarshal(xmlFile, &q)
}
