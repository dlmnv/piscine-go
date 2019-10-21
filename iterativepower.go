package piscine

func IterativePower(nb int, power int) int {
	a := 1

	for x := 1; x <= power; x ++ {
		a = a * nb 
	}
	return a
}

func main() {
	arg1 := 4
	arg2 := 3
	fmt.Println(IterativePower(arg1, arg2))
}
