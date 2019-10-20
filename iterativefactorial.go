package main

// import "fmt"

func IterativeFactorial(nb int) int {
	arg := 1

	for x := 1; x < nb; x++ {
		arg = arg * (x + 1)
	}
	return arg


}

// func main() {
// 	arg := 4
// 	fmt.Println(IterativeFactorial(arg))
// }
