package main

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/adrianratnapala/okpaaswoad"
)

func main() {
	if pw, err := okpaaswoad.ReadAndEncode(rand.Reader, 5); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", pw)
	}
}
