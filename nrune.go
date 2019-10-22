package piscine

func NRune(s string, n int) rune {
	if n > 0 && n <= slen {
		x := []rune(s)
		return x[n-1]
	}
	return 0
}

func slen(s string) {
	k := 0
	for range s {
		k++
	}
	return k
}
