package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for k := range arg {
		if k > 0 {
			arr := []rune(arg[k])
			for a := range arr {
				z01.PrintRune(arr[a])
			}
		}
		z01.PrintRune(10)
	}
}
