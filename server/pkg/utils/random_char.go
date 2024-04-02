package utils

import (
	"math/rand"
	"strings"
)

func RandomChar() string {
	var password string

	dic := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%&*()"
	split_string := strings.Split(dic, "")

	for i := 0; i < 21; i++ {
		r := rand.Intn(len(split_string))
		password += split_string[r]
	}

	return password
}
