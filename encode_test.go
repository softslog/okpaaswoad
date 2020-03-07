package okpaaswoad

import "testing"

func knownPw(super *testing.T, xpw string, bits ...byte) {
	super.Run(xpw, func(t *testing.T) {
		if pw := Encode(bits); pw != xpw {
			t.Errorf("%v encodes to %q != expected %q",
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
