package okpaaswoad

import "io"

// Encode encodes a byte array into password letters using Digraph().
// b should be random bits from some entropy source such as the crypto/rand
// package.
func Encode(b []byte) string {
	pw := make([]byte, 2*len(b))
	for k, bk := range b {
		pw[2*k], pw[2*k+1] = Digraph(bk)
	}
	return string(pw)
}

// Encode encodes bytes from an io.Reader into password letters using Digraph().
// r should be a source of random bits like Reader from package crypto/rand.
func ReadAndEncode(r io.Reader, n int) (string, error) {
	bits := make([]byte, n)
	if _, err := io.ReadFull(r, bits); err != nil {
		return "", err
	}

	return Encode(bits), nil
}
