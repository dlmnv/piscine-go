package piscine

func IsNumeric(str string) bool {
	slen := 0
	x := []rune(str)
	for range str {
		slen++
	}
	for i := 0; i < slen; i++ {
		if x[i] < '0' || x[i] > '9'
			return false
		}
	}
	return true
}
