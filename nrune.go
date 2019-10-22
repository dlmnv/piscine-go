package main

// import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	if n > 0 && n <= slen(s) {
		x := []rune(s)
		return x[n-1]
	}
	return 0
}

func slen(s string) int {
	k := 0
	for range s {
		k++
	}
	return k
}

// func main() {
// 	z01.PrintRune(NRune("Hello!", 3))
// 	z01.PrintRune(NRune("Salut!", 2))
// 	z01.PrintRune(NRune("Ola!", 4))
// 	z01.PrintRune('\n')
// }
