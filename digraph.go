package okpaaswoad

const (
	vowels     = "aei" + "ouy"
	consonants = "mnl" + "szr" + "bcdf" + "ghjkp" + "qtvwx"
)

// Digraph encodes a single byte b into two consecutive letters of a password.
// b should be random bits from some source such as the crypto/rand package.
// Both return values are lower-case English letters.
func Digraph(b byte) (byte, byte) {
	r128 := uint(b) >> 1
	cfbit := b & 1

	if r128 < 120 {
		d0 := vowels[r128/20]
		d1 := consonants[r128%20]
		if cfbit == 0 {
			return d0, d1
		} else {
			return d1, d0
		}
	}

	table := vowels
	if cfbit != 0 {
		table = consonants
	}

	r8 := r128 - 120
	return table[r8/3], table[r8%3+3]
}
