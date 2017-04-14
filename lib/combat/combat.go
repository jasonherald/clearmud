package combat

import (
	"fmt"
	"math/rand"
)

var (
	pHealth, mHealth, pDam, mDam int64
)

func Fight() int64 {
	pHealth = 100
	mHealth = 100
	for {

		pDam = rand.Int63n(20)
		mHealth = mHealth - pDam
		fmt.Println("Player hits mob for", pDam, ". Mob Health is now", mHealth)
		if mHealth <= 0 {
			break
		}

		mDam = rand.Int63n(20)
		pHealth = pHealth - mDam
		fmt.Println("Mob hits player for", pDam, ". Players Health is now", mHealth)
		if pHealth <= 0 {
			break
		}
	}
	if mHealth <= 0 {
		fmt.Println("Mob is dead, your Health is", pHealth)
	} else {
		fmt.Println("Player is dead")
	}
	return (pHealth)
}
