package okpaaswoad

import "io"

func Encode(bits []byte) string {
	pw := make([]byte, 2*len(bits))
	for k, b := range bits {
		pw[2*k], pw[2*k+1] = Digraph(b)
	}
	return string(pw)
}

func ReadAndEncode(r io.Reader, n int) (string, error) {
	bits := make([]byte, n)
	if nr, err := r.Read(bits); err == io.EOF {
		if nr < n {
			return "", io.ErrUnexpectedEOF
		}
	} else if err != nil {
		return "", err
	}

	return Encode(bits), nil
}
