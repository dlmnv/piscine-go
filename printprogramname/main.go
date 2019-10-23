package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	k := os.Args[0]
	arg := []rune(k)
	for x := range arg {
		z01.PrintRune(rune(arg[x]))
	}
	z01.PrintRune(10)
}
