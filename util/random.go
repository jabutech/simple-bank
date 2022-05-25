package util

import (
	"math/rand"
	"strings"
	"time"
)

// Const alphabet for use random data with string
const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Func init for first run
func init() {
	// Run rand.Seed
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	// Get total character on const alphabet
	k := len(alphabet)

	// Loop through n
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	// Use function RandomString for result 6 owner name
	return RandomString(6)
}

// RandomMonet generates a random amount of money
func RandomMoney() int64 {
	// Use function RandomInt for get number from 0 - 1000
	return RandomInt(0, 100)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	// Create list of currency
	currencies := []string{"IDR", "USD", "EUR"}
	// Get len currency
	n := len(currencies)
	// Return currency with rand much as length currencies
	return currencies[rand.Intn(n)]
}
