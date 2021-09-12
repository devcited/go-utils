package random

import (
	crypto_rand "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/rand"
)

func Init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}

	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

// Seed ...
func Seed(value int64) {
	rand.Seed(value)
}

// GenerateID ...
func GenerateID() string {
	return GenerateString(12)
}

// GenerateString ...
func GenerateString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}
