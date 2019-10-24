package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(str string) {
	for _, s := range str {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}

func SortString(table []string) {

	len := 0
	for i := range table {
		len = i
	}

	count := 1
	for count != 0 {
		count = 0
		for i, j := 1, 2; j <= len; i, j = i+1, j+1 {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
				count++
			}
		}
	}
}

func main() {
	args := os.Args

	SortString(args)

	for i, str := range args {
		if i > 0 {
			PrintStr(str)
		}
	}
}
