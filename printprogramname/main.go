package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arg := []rune(os.Args[0])
	for _, i := range arg {
		z01.PrintRune(rune(arg[i]))
	}
	z01.PrintRune('\n')
}
