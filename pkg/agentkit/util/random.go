package util

import (
	"math/rand"
	"time"
)

// RandomIntN returns a random integer between 0 and `n`.
// (This function takes care of seeding the random generator, too.)
func RandomIntN(n int) int {
	rand.Seed(time.Now().UnixNano() * 7)
	return rand.Intn(n)
}
