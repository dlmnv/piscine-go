package main

import "github.com/01-edu/z01"

func PrintComb() {
	for a := 0; a <=7; a++ {
		for b := 1; b <=8; b++ {
			for c := 2; c <= 9; c++ {
					if a := 7; b := 8; c := 9{
						z01.PrintComb("")
					}else{
						z01.PrintComb(",")
					}
					z01.PrintComb(a)
					z01.PrintComb(b)
					z01.PrintComb(c)

				}
			}
		}
	}
}

func main() {
	z01.PrintComb()
}