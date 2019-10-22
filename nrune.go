package piscine

func NRune(s string, n int) rune {
	x := []rune(s)
	return x[n-1]
}