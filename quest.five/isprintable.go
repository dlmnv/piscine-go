package piscine

func IsPrintable(str string) bool {
	slen := 0
	x := []rune(str)
	for range str {
		slen++
	}
	for i := 0; i < slen; i++ {
		if x[i] < 32 || x[i] > 127 {
			return false
		}
	}
	return true
}
