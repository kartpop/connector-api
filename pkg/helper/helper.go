package helper

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetMD5Hash(text string) string {
	// to ensure unique hash in case text is duplicate, prepend random string to text
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	// make hash
	hash := md5.Sum([]byte(string(b) + text))
	return hex.EncodeToString(hash[:])
}
