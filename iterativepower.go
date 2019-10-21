package piscine

func IterativePower(nb int, power int) int {
	a := 1

	for x := 1; x <= power; x++ {
		a = a * nb 
	}
	return a
}
