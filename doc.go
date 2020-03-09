/*
Package okpaaswoad generates random, humane passwords (okpaaswoads) by
encoding bits of entropy into pairs of lower-case English letters called
digraphs.  These are chosen so that the resulting passwords can be pronounced
out loud.  Okpaaswoads are secure in the sense that they reversibly encode all
the bits of entropy.

API

Function Encode() encodes bytes and returns an okpaaswoad as a string. It
can't return an error.  Function ReadAndEncode() encodes a given number of
bytes from an io.Reader that provides entropy, for example:

	pw, err := okpaaswoard.ReadAndEncode(rand.Reader, bytesOfEntropy)

There is no default entropy source, most users should use package
"crypto/rand".

Function Digraph() encodes a single byte into a pair of letters.  There is no
facility for decoding okpaaswoads or digraphs into bytes -- users who want
this can use Digraph() to build a 26x26 lookup table.
*/
package okpaaswoad
