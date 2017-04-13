package ticker

import (
	"fmt"
	"time"
)

var ticker = time.NewTicker(time.Second * 5)

func Start() {
	go func() {
		for t := range ticker.C {
			// execute some stuff
			fmt.Println("Tick at", t)
		}
	}()
}

func Stop() {
	ticker.Stop()
}
