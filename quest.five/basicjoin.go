package piscine

func BasicJoin(strs []string) string {
	slen := 0
	var k string
	for range strs {
		slen++
	}
	for value := 0; value < slen; value++ {
		k = k + strs[value]
	}
	return k
}
