package impl

import (
	"crypto/sha256"
	"fmt"
)

type Sha256Raw struct{}

func (s Sha256Raw) GetName() string {
	return "SHA256 Raw"
}

func (s Sha256Raw) Generate(password string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
}

func (s Sha256Raw) Validate(password, hash string) bool {
	return s.Generate(password) == hash
}
