package helpers

import (
	"fmt"
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

func GenerateSessionId() string {
	rand.Seed(time.Now().Unix())
	min := 1000
	max := 9999
	return fmt.Sprintf("sessionId%v", rand.Intn(max-min+1)+min)
}

func GenerateProcessId() string {
	rand.Seed(time.Now().Unix())
	min := 1000
	max := 9999
	return fmt.Sprintf("processId%v", rand.Intn(max-min+1)+min)
}
