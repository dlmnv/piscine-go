package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	k := os.Args[0]
	arg := []rune(k)
	z01.PrintRune(arg[0])
}
