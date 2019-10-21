package main

import "fmt"

func IterativeFactorial(nb int) int {
	if nb >= 0 && nb <= 20 {
		arg := 1

		for x := 1; x < nb; x++ {
		arg = arg * (x + 1)
		}
		return arg
	}
	return 0
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
