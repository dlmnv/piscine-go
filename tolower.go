package piscine

func ToLower(s string) string {
	slen := 0
	k := []rune(s)
	for range s {
		slen++
	}
	for value := 0; value < slen; value++ {
		if k[value] >= 'A' && k[value] <= 'Z' {
			k[value] = (k[value] + 32)
		}
	}
	lower := string(k)
	return lower
}
