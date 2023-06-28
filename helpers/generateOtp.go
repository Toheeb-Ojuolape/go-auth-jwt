package helpers

import (
	"math/rand"
	"time"
)

// write code to generate a 6-digit otp that will be sent for verification purposes
func GenerateOtp() int {
	rand.Seed(time.Now().Unix())
	min := 100000
	max := 999999
	return rand.Intn(max-min+1) + min
}
