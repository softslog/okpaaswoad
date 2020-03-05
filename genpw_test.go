package okpasswoad

import (
	"fmt"
	"testing"
)

type log interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

func r8Decode(idx0, idx1 uint) (byte, error) {
	if idx0 >= 3 {
		return 0, fmt.Errorf("Illegal first index %d in r8 pair", idx0)
	}
	if idx1 < 3 || idx1 >= 6 {
		return 0, fmt.Errorf("Illegal second index %d in r8 pair", idx1)
	}

	if r8 := 3*idx0 + (idx1 - 3); r8 >= 8 {
		return byte(r8), fmt.Errorf("Illegal r8 combo = %d <- (%d, %d)",
			r8, idx0, idx1)
	} else {
		return byte(r8), nil
	}
}

func decodeDoubleVowel(idx0, idx1 uint) (byte, error) {
	if int(idx0) >= len(vowels) {
		panic(fmt.Sprintf("Bad vowel-index idx0=%d", idx0))
	}
	if int(idx1) >= len(vowels) {
		panic(fmt.Sprintf("Bad vowel-index idx1=%d", idx1))
	}

	// Return the r8Decode (without the consonst-first bit).
	if r8, err := r8Decode(idx0, idx1); err != nil {
		return 0, fmt.Errorf(
			"Vowel pair %d(%c), %d(%c) can't r8-combine: %v",
			idx0, vowels[idx0],
			idx1, vowels[idx1],
			err)
	} else {
		return byte(120+r8) * 2, nil
	}
}

// findConsonant Searches the consonant list for `d` and returns its index or
// an error.  It is assumed that `d` is already provably a non-vowel.
func findConsonant(d byte) (uint, error) {
	for idx := 0; idx < len(consonants); idx++ {
		if consonants[idx] == d {
			return uint(idx), nil
		}
	}

	return 0, fmt.Errorf("%q(0x%x) is neither vowel nor a consonant", d, d)
}

func decodeVowCon(vidx, cidx uint) byte {
	if int(vidx) >= 6 {
		panic(fmt.Sprintf("Bad vowel-index vidx=%d", vidx))
	}
	if int(cidx) >= 20 {
		panic(fmt.Sprintf("Bad consonant-index cidx=%d", cidx))
	}

	r := vidx*20 + cidx
	if r >= 128 {
		panic(fmt.Sprintf("decodeVowCon(%d, %d) = %d > 127", vidx, cidx, r))
	}
	return byte(r)
}

func decodeVowelFirst(idx0 uint, d1 byte) (byte, error) {
	for idx1 := 0; idx1 < len(vowels); idx1++ {
		if vowels[idx1] == d1 {
			return decodeDoubleVowel(idx0, uint(idx1))
		}
	}

	if idx1, err := findConsonant(d1); err != nil {
		return 0, err
	} else {
		return 2 * decodeVowCon(idx0, idx1), nil
	}
}

func decodeDigraph(d0, d1 byte, l log) (byte, error) {
	// If d0 is a vowel, hand off to decodeVowelFirst.
	for idx0 := 0; idx0 < len(vowels); idx0++ {
		if vowels[idx0] == d0 {
			l.Logf("%q[%d] is a vowel, decodeVowelFirst d1=%d",
				d0, idx0, d1)
			return decodeVowelFirst(uint(idx0), d1)
		}
	}

	// d0 is not a vowel, insist it is a consonant.
	idx0, err := findConsonant(d0)
	if err != nil {
		return 0, err
	}

	// If d1 is a vowel decodeVowCon it, then set the consonant-first bit.
	for idx1 := 0; idx1 < len(vowels); idx1++ {
		if vowels[idx1] == d1 {
			return 2*decodeVowCon(uint(idx1), idx0) + 1, nil
		}
	}

	// d0 is not a vowel, insist it is a consonant.
	idx1, err := findConsonant(d1)
	if err != nil {
		return 0, err
	}

	// r8Decode but set the consonant-first bit
	if r8, err := r8Decode(idx0, idx1); err != nil {
		return r8, fmt.Errorf(
			"Consonant pair %d(%c), %d(%c) can't r8-combine: %v",
			idx0, consonants[idx0],
			idx1, consonants[idx1],
			err)
	} else {
		return byte(120+r8)*2 + 1, nil
	}
}

func isSmallLetter(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func TestDigraph(t *testing.T) {
	for k := 0; k < 256; k++ {
		bits := byte(k)
		d0, d1 := Digraph(bits)
		if !isSmallLetter(d0) {
			t.Errorf("Digraph d0=%q is not a small letter", d0)
		}
		if !isSmallLetter(d1) {
			t.Errorf("Digraph d1=%q is not a small letter", d1)
		}
		t.Logf("Attempting x%x == decodeDigraph(%q, %q)", bits, d0, d1)
		roundTrip, err := decodeDigraph(d0, d1, t)
		if err != nil {
			t.Errorf("decodeDigraph(%x => %d, %d) failed: %v ",
				bits, d0, d1, err)
		} else {
			if int(roundTrip) != k {
				t.Errorf("Digraph(%d) = (%c, %c) decodes to %d.",
					bits, d0, d1, roundTrip)
			} else {
				t.Logf("GOOD x%x == decodeDigraph(%q, %q)", bits, d0, d1)
			}
		}
	}
}
