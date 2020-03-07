package okpaaswoad

func Encode(bits []byte) string {
	pw := make([]byte, 2*len(bits))
	for k, b := range bits {
		pw[2*k], pw[2*k+1] = Digraph(b)
	}
	return string(pw)
}
