package piscine

func LastRune(s string) rune {
	x := []rune(s)
	slen := 0
	for range s {
		slen++
	}
	return x[slen-1]
}
