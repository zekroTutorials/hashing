package impl

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type Sha256Salted struct{}

func (s Sha256Salted) GetName() string {
	return "SHA256 Salted"
}

func (s Sha256Salted) Generate(password string) string {
	salt := make([]byte, 8)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}

	return s.sum(password, salt)
}

func (s Sha256Salted) Validate(password, hash string) bool {
	iSalt := strings.LastIndex(hash, "$") + 1
	salt, err := hex.DecodeString(hash[iSalt:])
	if err != nil {
		panic(err)
	}

	return s.sum(password, salt) == hash
}

func (s Sha256Salted) sum(password string, salt []byte) string {
	buff := []byte(password)
	buff = append(buff, salt...)
	return fmt.Sprintf("%x$%x", sha256.Sum256(buff), salt)
}
