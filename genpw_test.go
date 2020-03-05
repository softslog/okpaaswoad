package okpasswoad

import "testing"

func decodeDigraph(d0, d1 byte) byte {
	return d0
}

func TestDigraph(t *testing.T) {
	for k := 0; k < 256; k++ {
		d0, d1 := Digraph(byte(k))
		if roundTrip := decodeDigraph(d0, d1); int(roundTrip) != k {
			t.Errorf("Digraph(%d) = (%c, %c) decodes to %d.",
				k, d0, d1, roundTrip)
		}
	}
}
