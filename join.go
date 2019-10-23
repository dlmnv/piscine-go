package piscine

func Join(strs []string, sep string) string {
	slen := 0
	var k string
	for range strs {
		slen++
	}
	for value := 0; value < slen; value++ {
		if value == slen-1 {
			k = k + strs[value]
		} else {
		k = k + strs[value] + sep
	}
	return k
}
