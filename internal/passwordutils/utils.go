package passwordutils

import (
	"math/rand"
	"time"
)

func ObtainRandomSeed() {
	rand.Seed(time.Now().UnixNano())
}
