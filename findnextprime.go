package piscine

func FindNextPrime(nb int) int {
	if nb > 1 {
		for x := 2; x < nb; x++ {
			if nb%x == 0 {
				return FindNextPrime(nb+1)
			}
		}
	}
	return nb
}