package impl

import (
	"runtime"

	"github.com/alexedwards/argon2id"
)

type Sha256Argon2ID struct{}

var argon2IDParams = &argon2id.Params{
	Memory:      128 * 1024,
	Iterations:  4,
	Parallelism: uint8(runtime.NumCPU()),
	SaltLength:  16,
	KeyLength:   32,
}

func (s Sha256Argon2ID) Generate(password string) string {
	hash, err := argon2id.CreateHash(password, argon2IDParams)
	if err != nil {
		panic(err)
	}

	return hash
}

func (s Sha256Argon2ID) Validate(password, hash string) bool {
	match, _, err := argon2id.CheckHash(password, hash)
	if err != nil {
		panic(err)
	}

	return match
}
