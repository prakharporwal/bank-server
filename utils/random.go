package utils

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz.+")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomEmail() string {
	name := RandomString(10)
	company := RandomString(5)

	email := strings.Builder{}
	email.WriteString(name)
	email.WriteString("@")
	email.WriteString(company)
	email.WriteString(".com")

	return email.String()
}

func RandomNumber64() int64 {
	return rand.Int63() + 1 // adding 1 will never make it zero
}

func RandomNumber() int {
	return rand.Int() + 1 // adding 1 will never make it zero
}

func RandomCurrency() string {
	currency := []string{"CAD", "INR", "USD"}
	n := len(currency)
	return currency[rand.Intn(n)]
}
