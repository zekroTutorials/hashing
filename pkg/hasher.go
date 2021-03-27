package pkg

type Hasher interface {
	Generate(password string) string
	Validate(password, hash string) bool
}
