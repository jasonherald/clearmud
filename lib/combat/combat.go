package combat

import (
	"math/rand"
)

func Fight() int {
	pHealth := 100
	mHealth := 100
	for {

		pDam := rand.Intn(20)
		mHealth = mHealth - pDam
		//	fmt.Println("Player hits mob for", pDam, ". Mob Health is now", mHealth)
		if mHealth <= 0 {
			break
		}

		mDam := rand.Intn(20)
		pHealth = pHealth - mDam
		//	fmt.Println("Mob hits player for", pDam, ". Players Health is now", mHealth)
		if pHealth <= 0 {
			break
		}
		//	}
		//	if mHealth <= 0{
		//	fmt.Println("Mob is dead, your Health is", pHealth)
		//	} else {
		//	fmt.Println("Player is dead")
	}
	return (pHealth)
}
