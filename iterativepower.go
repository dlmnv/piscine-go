package main

import "fmt"

func IterativePower(nb int, power int) int {
	if power > 0 {
		x := 1
		for a := 1; a <= power; a++ {
			x = x * nb
		}
		return x
	} else if power == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	arg1 := 4
	arg2 := 3
	fmt.Println(IterativePower(arg1, arg2))
}