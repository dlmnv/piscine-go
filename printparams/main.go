package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	arr := []rune(os.Args)
	for k := range arg {
		for k > 0 {
			for a := range arr {
				z01.PrintRune('a')
			}
		}
		z01.PrintRune(10)
	}
}
