package pubtools

import (
	"math/rand"
	"time"
)

// RandomInt returns random int between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomDuration returns random duration between min and max
func RandomDuration(min, max int) time.Duration {
	return time.Duration(RandomInt(min, max))
}
