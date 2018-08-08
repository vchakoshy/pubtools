package pubtools

import (
	"math/rand"
	"time"
)

// RandomLib struct
type RandomLib struct{}

// Random returns RandomLib
func Random() *RandomLib {
	return &RandomLib{}
}

// Int returns random int between min and max
func (r *RandomLib) Int(min, max int) int {
	return rand.Intn(max-min) + min
}

// Duration returns random duration between min and max
func (r *RandomLib) Duration(min, max int) time.Duration {
	return time.Duration(r.Int(min, max))
}
