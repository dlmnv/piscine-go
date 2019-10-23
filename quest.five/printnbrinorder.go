package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	a := 0
	array := []rune{}
	z := 0

	if n == 0 {
		array = append(array, rune(n))
	}

	for n != 0 {
		a = n % 10
		n = n / 10
		array = append(array, rune(a))
	}

	for range array {
		z++
	}

	for i := 0; i < z-1; i++ {
		for j := 0; j < z-1-i; j++ {
			if array[j] > array[j+1] {
				k := array[j]
				array[j] = array[j+1]
				array[j+1] = k
			}
		}
	}

	for i := 0; i < z; i++ {
		z01.PrintRune(array[i] + 48)
	}
	// z01.PrintRune('\n')
}
