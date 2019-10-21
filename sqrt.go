package piscine

func Sqrt(nb int) int {
	k := 0
	for x := 1; x <= nb; x++ {
		if x*x == nb {
			k = x
		}
	}
	return k
}
