package okpaaswoad

import (
	"bytes"
	"io"
	"testing"
)

func knownPw(super *testing.T, xpw string, bits ...byte) {
	super.Run("Encode(xpw)", func(t *testing.T) {
		if pw := Encode(bits); pw != xpw {
			t.Errorf("%q Encode()s to %q != expected %q",
				bits, pw, xpw)
		}
	})

	super.Run("ReadAndEncode(xpw)", func(t *testing.T) {
		longbits := append(bits, []byte("rubbish")...)
		r := bytes.NewReader(bits)
		if pw, err := ReadAndEncode(r, len(bits)); err != nil {
			t.Errorf("%q fails to ReadeAndEncode(): %v", longbits, err)
		} else if pw != xpw {
			t.Errorf("%q ReadeAndEncode()s to %q != expected %q",
				bits, pw, xpw)
		}
	})
}

func TestKnownPasswords(t *testing.T) {
	knownPw(t, "")
	knownPw(t, "okpaaswoad", 146, 29, 6, 157, 16)
	knownPw(t, "amma", 00, 01)
	knownPw(t, "aoauayeoeueyioiu", 0xf0, 0xf2, 0xf4, 0xf6, 0xf8, 0xfa, 0xfc, 0xfe)
	knownPw(t, "msmzmrnsnznrlslz", 0xf1, 0xf3, 0xf5, 0xf7, 0xf9, 0xfb, 0xfd, 0xff)
}

func TestShortRead(t *testing.T) {
	bits := []byte("short")
	nentropy := len(bits) + 1
	r := bytes.NewReader(bits)
	if _, err := ReadAndEncode(r, nentropy); err != io.ErrUnexpectedEOF {
		t.Errorf("reading %d bytes from %q should give "+
			"io.ErrUnexpectedEOF, not %v",
			nentropy, bits, err)
	}
}
