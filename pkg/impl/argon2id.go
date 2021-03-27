package impl

import (
	"runtime"

	"github.com/alexedwards/argon2id"
)

type Argon2id struct{}

var argon2IDParams = &argon2id.Params{
	Memory:      128 * 1024,
	Iterations:  4,
	Parallelism: uint8(runtime.NumCPU()),
	SaltLength:  16,
	KeyLength:   32,
}

func (s Argon2id) GetName() string {
	return "Argon2id"
}

func (s Argon2id) Generate(password string) string {
	hash, err := argon2id.CreateHash(password, argon2IDParams)
	if err != nil {
		panic(err)
	}

	return hash
}

func (s Argon2id) Validate(password, hash string) bool {
	match, _, err := argon2id.CheckHash(password, hash)
	if err != nil {
		panic(err)
	}

	return match
}
