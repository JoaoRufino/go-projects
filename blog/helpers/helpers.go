package helpers

import (
	"crypto/rand"
	"fmt"
)

func GenerateId() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
