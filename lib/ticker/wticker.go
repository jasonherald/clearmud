package ticker

import (
	"time"
)

var ticker = time.NewTicker(time.Second * 5)

//Start the ticker
func Start() {
	go func() {
		for _ = range ticker.C {
			UpdateTime()
		}
	}()
}

//Stop the ticker
func Stop() {
	ticker.Stop()
}
