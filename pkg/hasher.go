package pkg

// Hasher provides general functionalities to
// generate a hash from a given string and
// validate a given string by comparing it to
// the passed hash.
type Hasher interface {
	Generate(password string) string
	Validate(password, hash string) bool
}
