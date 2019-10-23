package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	len := 0
	for range arg {
		len++
	}
	for indexSlova := len - 1; indexSlova > 0; indexSlova-- {
		massive := []rune(arg[indexSlova])
		for indexBukvi := range massive {
			z01.PrintRune(massive[indexBukvi])
		}
		z01.PrintRune(10)
	}
}
