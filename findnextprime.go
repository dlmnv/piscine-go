package piscine

func FindNextPrime(nb int) int {
	for x := 2; x <= nb; x++ {
		if nb%x != 0 {
			return nb
		} else {
			return FindNextPrime(nb + 1)
		}
	}
	return nb
}