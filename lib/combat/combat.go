package combat

import (
	"math/rand"
)

var (
	pHealth, mHealth, pDam int64
)

func Fight() int64 {
	pHealth = 100
	mHealth = 100

	pDam = rand.Int63n(20)
	mHealth = mHealth - pDam
	return (pHealth)
}
