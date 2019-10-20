package piscine

func IterativeFactorial(nb int) int {
	arg := 1

	for x := 1; x < nb; x++ {
		arg = arg * (x + 1)
	}
	return arg
}
