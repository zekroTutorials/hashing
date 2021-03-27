package impl

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

var pepperStore = func(n int) (p [][]byte) {
	p = make([][]byte, n)
	for i := range p {
		p[i] = make([]byte, 8)
		if _, err := rand.Read(p[i]); err != nil {
			panic(err)
		}
	}
	return
}(2000)

type Sha256Peppered struct{}

func (s Sha256Peppered) GetName() string {
	return "SHA256 Peppered"
}

func (s Sha256Peppered) Generate(password string) string {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(len(pepperStore))))
	if err != nil {
		panic(err)
	}

	return s.sum(password, pepperStore[i.Int64()])
}

func (s Sha256Peppered) Validate(password, hash string) bool {
	for _, p := range pepperStore {
		if s.sum(password, p) == hash {
			return true
		}
	}
	return false
}

func (s Sha256Peppered) sum(password string, pepper []byte) string {
	buff := []byte(password)
	buff = append(buff, pepper...)
	return fmt.Sprintf("%x", sha256.Sum256(buff))
}
