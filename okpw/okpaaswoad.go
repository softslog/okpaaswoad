package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ratnapala/okpaaswoad"
)

var (
	dgTable = flag.Bool("dg-table", false, "Print all possible digraph encodings.")
	okOrder = flag.Bool("okorder", false,
		"Actions (such as -dg-table) display okpaaswoard order.")
)

func printDigraphTable(w io.Writer, permute func(int) int) {
	var rmap [26][26]uint16
	for k := 0; k < 256; k++ {
		d0, d1 := okpaaswoad.Digraph(byte(k))
		rmap[d0-'a'][d1-'a'] = uint16(k + 1)
	}

	ktitle := "                                SECOND LETTER                                "
	//         00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f 10 11 12 13 14 15 16 17 18 19
	jtitle := "      FIRST LETTER        "

	fmt.Fprintf(w, "    %s\n    ", ktitle)
	for k := 0; k < 26; k++ {
		fmt.Fprintf(w, " %c ", 'a'+permute(k))
	}
	w.Write([]byte("\n    "))
	for k := 0; k < 26; k++ {
		w.Write([]byte("-+-"))
	}
	w.Write([]byte("\n"))

	for j := 0; j < 26; j++ {
		lj := permute(j)
		fmt.Fprintf(w, "%c %c|", jtitle[j], 'a'+lj)

		for k := 0; k < 26; k++ {
			lk := permute(k)
			if r := rmap[lj][lk]; r == 0 {
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

	letterOrder := func(k int) int { return k }
	if *okOrder {
		letterOrder = func(k int) int {
			return int("aeiouymnlszrbcdfghjkpqtvwx"[k] - 'a')
		}
	}

	if *dgTable {
		printDigraphTable(os.Stdout, letterOrder)
		return
	}

	pw := okpaaswoad.Encode([]byte{146, 29, 6, 157, 16})
	fmt.Printf("%s\n", pw)
}
