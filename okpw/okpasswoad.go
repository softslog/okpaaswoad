package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	okpasswoad "github.com/ratnapala/okpaaswoad"
)

var (
	dgTable = flag.Bool("dg-table", false, "Print all possible digraph encodings")
)

func printDigraphTable(w io.Writer) {
	var rmap [26][26]uint16
	for k := 0; k < 256; k++ {
		d0, d1 := okpasswoad.Digraph(byte(k))
		rmap[d0-'a'][d1-'a'] = uint16(k + 1)
		log.Printf("%c, %c = Digraph(%d)", d0, d1, k)
	}

	ktitle := "                                second letter                                "
	//         00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 10 11 12 13 14 15 16 17 18 19
	jtitle := "      first letter        "

	fmt.Fprintf(w, "    %s\n   ", ktitle)
	for k := 0; k < 26; k++ {
		fmt.Fprintf(w, " %c ", 'a'+k)
	}
	w.Write([]byte{'\n'})

	for j := 0; j < 26; j++ {
		fmt.Fprintf(w, "%c %c", jtitle[j], 'a'+j)

		for k := 0; k < 26; k++ {
			if r := rmap[j][k]; r == 0 {
				fmt.Fprintf(w, " . ")
			} else {
				fmt.Fprintf(w, " %02x", r-1)
			}
		}
		w.Write([]byte{'\n'})
	}
}

func main() {
	flag.Parse()

	if *dgTable {
		printDigraphTable(os.Stdout)
		return
	}

	for _, b := range []byte{146, 29, 6, 157, 16} {
		d0, d1 := okpasswoad.Digraph(b)
		fmt.Printf("%c%c\n", d0, d1)
	}
}
