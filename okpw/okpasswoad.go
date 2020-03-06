package main

import (
	"fmt"

	okpasswoad "github.com/ratnapala/okpaaswoad"
)

func main() {
	for _, b := range []byte{146, 29, 6, 157, 16} {
		d0, d1 := okpasswoad.Digraph(b)
		fmt.Printf("%c%c\n", d0, d1)
	}
}
