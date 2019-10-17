package piscine

import "fmt"

func Swap(a *int, b *int) {
	var x int
	var z int
	x = 1
	z = 0
	*a = x
	*b = z
}

func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
