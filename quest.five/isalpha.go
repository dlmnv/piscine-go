package piscine

func IsAlpha(str string) bool {
	slen := 0
	x := []rune(str)
	for range str {
		slen++
	}
	for i := 0; i < slen; i++ {
		if x[i] < '0' || (x[i] > '9' && x[i] < 'A') || (x[i] > 'Z' && x[i] < 'a') || x[i] > 'z' {
			return false
		}
	}
	return true
}
