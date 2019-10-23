package piscine

func IsUpper(str string) bool {
	slen := 0
	x := []rune(str)
	for range str {
		slen++
	}
	for i := 0; i < slen; i++ {
		if x[i] < 'A' || x[i] > 'Z' {
			return false
		}
	}
	return true
}
