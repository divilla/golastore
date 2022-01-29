package random

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// Charsets.
const (
	Uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase    = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic   = Uppercase + Lowercase
	Numeric      = "0123456789"
	Alphanumeric = Alphabetic + Numeric
	Symbols      = "`" + `~!@#$%^&*()-_+={}[]|\;:"<>,./?`
	URI3986Safe  = `-_.~`
	URI2396Safe  = `-_.!~*'()`
	Hex          = Numeric + "abcdef"
)

// MustString builds cryptographically safe string of specified length & using specified charsets
// returning result or panicking on error.
func MustString(length int, charsets ...string) string {
	out, err := String(length, charsets...)
	if err != nil {
		panic(err)
	}

	return out
}

// MustURINew builds cryptographically safe, URI safe (according to RFC3985) string
// of specified length & using specified charsets returning result or panicking on error.
func MustURINew(length int) string {
	out, err := String(length, Alphanumeric, URI3986Safe)
	if err != nil {
		panic(err)
	}

	return out
}

// MustURI builds cryptographically safe, old URI safe (according to RFC2396) string
// of specified length & using specified charsets returning result or panicking on error.
func MustURI(length int) string {
	out, err := String(length, Alphanumeric, URI2396Safe)
	if err != nil {
		panic(err)
	}

	return out
}

// String builds cryptographically safe string of specified length & using specified charsets
// returning result or error.
func String(length int, charsets ...string) (string, error) {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = Alphanumeric
	}
	l := int64(len(charset)) - 1

	var sb strings.Builder
	for i := 0; i < length; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(l))
		if err != nil {
			return "", err
		}
		n := nBig.Int64()
		sb.WriteString(charset[n : n+1])
	}

	return sb.String(), nil
}
