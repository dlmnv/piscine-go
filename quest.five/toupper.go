package piscine

func ToUpper(s string) string {
	slen := 0
	k := []rune(s)
	for range s {
		slen++
	}
	for value := 0; value < slen; value++ {
		if k[value] >= 'a' && k[value] <= 'z' {
			k[value] = (k[value] - 32)
		}
	}
	upper := string(k)
	return upper
}
