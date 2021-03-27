package main

import (
	"fmt"
	"time"

	"github.com/zekroTutorials/hashing/pkg"
	"github.com/zekroTutorials/hashing/pkg/impl"
)

var (
	passwords = []string{
		"pw1",
		"pw1",
		"my_password",
		"123456",
		"123456",
		"My_5up3r_5t0ng_PASSWORD_!",
	}
)

type pair struct {
	Password string
	Hash     string
}

func more(i, v int) int {
	if v > i {
		return v
	}
	return i
}

func timeit(action func()) time.Duration {
	start := time.Now()
	action()
	return time.Since(start)
}

func computeHashes(hasher pkg.Hasher) {
	var maxLenPw, maxLenHash int
	pairs := make([]pair, len(passwords))

	tookForHasing := timeit(func() {
		for i, pw := range passwords {
			hash := hasher.Generate(pw)
			pairs[i].Password = pw
			pairs[i].Hash = hash
			maxLenPw = more(maxLenPw, len(pw))
			maxLenHash = more(maxLenHash, len(hash))
		}
	})

	tookForValidation := timeit(func() {
		for _, p := range pairs {
			if !hasher.Validate(p.Password, p.Hash) {
				panic("validation failed")
			}
		}
	})

	format := fmt.Sprintf("%%-%ds - %%%ds\n", maxLenPw, maxLenHash)
	for _, p := range pairs {
		fmt.Printf(format, p.Password, p.Hash)
	}

	fmt.Printf(
		"\nTime for Hashing:    %s\n"+
			"Time for Validation: %s\n",
		tookForHasing.String(),
		tookForValidation.String())
}

func main() {
	hashers := []pkg.Hasher{
		// impl.Sha256Raw{},
		// impl.Sha256Salted{},
		// impl.Sha256Peppered{},
		impl.Sha256Argon2ID{},
	}

	for _, hasher := range hashers {
		computeHashes(hasher)
	}
}