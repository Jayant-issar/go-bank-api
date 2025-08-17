package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "qwertyuioplkjhgfdsazxcvbnm"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) //min-max

}

// Random string of n characters
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		char := alphabet[rand.Intn(k)]
		sb.WriteByte(char)
	}

	return sb.String()
}

// Random owner name
func RandomOwner() string {
	return RandomString(6) //generate a random owner name with 6 letters in it
}

func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD", "INR"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
