package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := []rune(os.Args[0])
	for i := range arg {
		z01.PrintRune(arg[i])
	}
	z01.PrintRune('\n')
}
