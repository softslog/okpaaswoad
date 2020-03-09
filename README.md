Okpaaswoad generates random, humane passwords.
==============================================

Author     : Adrian Ratnapala
Get Source : git clone https://github.com/adrianratnapala/okpaaswoad.gitA
More info  : https://github.com/adrianratnapala/

Copyright (C) 2020, Adrian Ratnapala, under the ISC license. See file LICENSE.


Okpaaswoad is a small tool and Go library (module) for generating strong
random passwords, such as:

        vihaleqaco
        ryompuibyv
        ubqasojywi
        ocyftoaqfu

Passwords like these are secure in the sense that they encode 40 bits of
entropy.  But they are easier to handle than most random passwords because:

* They can be a spoken out loud, which is an aid to memory.
* Being reasonably short, and all lower-case, they are easy to type.

The Tool
--------

Go users can install the `okpw` tool with

        go get github.com/adrianratnapala/okpaaswoad/okpw
        go install github.com/adrianratnapala/okpaaswoad/okpw

Then you can just run it to write random passwords to standard output.

        $ okpw
        fozykucuac

By default `okpw` gets entropy from the `crypto/rand` Go package. You can
choose any other source with the `-entropy` flag:

        $ okpw -entropy=filaname-of-entropy-source

And `-` means standard input.  Thus:

        $ okpw -entropy=filaname-of-entropy-source -entropy=- < /dev/urandom

and:

        $ okpw -entropy=filaname-of-entropy-source -entropy=/dev/urandom

do the same thing.

The Library
------------

Install it with

        go get github.com/adrianratnapala/okpaaswoad

You can then generate passwords with

        pw, err := okpaaswoad.ReadAndEncode(rand.Reader, bytesOfEntropy)

See also [doc/example.go](doc/example.go) and refer to the package
documentation for further details.

        go doc github.com/adrianratnapala/okpaaswoad
