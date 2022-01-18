package security

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// GenerateRandomBytes returns securely generated random bytes
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded securely generated random string.
func GenerateRandomString(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

// GenerateRandomCode returns a securely generated random string that consists of numbers that has length of 2*n
func GenerateRandomCode(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}

	code := ""
	for i := 0; i < n; i++ {
		code = code + fmt.Sprintf("%02d", b[i]%100)
	}

	return code, err
}

// GenerateRandomStringByPool returns a securely generated random string based on pool.
func GenerateRandomStringByPool(n int, pool string) (string, error) {
	l := byte(len(pool))

	b, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}

	for i := 0; i < n; i++ {
		b[i] = pool[(b[i])%l]
	}

	return string(b), nil
}
