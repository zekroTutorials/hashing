# Password Hashing Concepts

This repository holds some example implementations of general concepts used in password hashing.

> **ATTENTION:**  
> The shown exaples are **not** production ready implementations and shall not be used for any **real** password security implementation without further security fortification!

## Index

The hasher implementations below implement the [hasher](pkg/hasher.go) interface.
```go
type Hasher interface {
	Generate(password string) string
	Validate(password, hash string) bool
}
```

- [SHA256 Raw](pkg/impl/sha256_raw.go)
- [SHA256 Salted](pkg/impl/sha256_salted.go)
- [SHA256 Peppered](pkg/impl/sha256_peppered.go)
- [Argon2id](pkg/impl/argon2id.go)

## Resources

This repository was created as demonstration for my [*(german)* introduction video](https://youtu.be/OcKlnvKXbpg) about password hashing.

Here you can find more resources *(in english)* about the topic:

### Videos

- https://youtu.be/8ZtInClXe1Q
- https://youtu.be/b4b8ktEV4Bg
- https://youtu.be/7U-RbOKanYs

### Information about Hashing Concepts

- https://auth0.com/blog/hashing-passwords-one-way-road-to-security/
- https://www.gearbrain.com/password-security-hashing-salting-peppering-2647766220.html
- https://nordpass.com/blog/pepper-password/

### Information about Argon2

- https://en.wikipedia.org/wiki/Argon2
- https://www.twelve21.io/how-to-choose-the-right-parameters-for-argon2/
- https://github.com/P-H-C/phc-winner-argon2

### Useful Online Hasing Tools

- https://md5hashing.net/hash
- https://crackstation.net/
- https://www.devglan.com/online-tools/rsa-encryption-decryption

---

© 2020 Ringo Hofffmann (zekro Development).  
Covered by the MIT License.