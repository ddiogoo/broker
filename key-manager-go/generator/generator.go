package generator

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GenerateApiKey generate an API key based on UUID.
func GenerateApiKey() string {
	h, err := bcrypt.GenerateFromPassword([]byte(uuid.New().String()), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	hasher := md5.New()
	hasher.Write(h)
	return hex.EncodeToString(hasher.Sum(nil))
}
