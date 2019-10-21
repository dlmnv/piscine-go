package main

// import "fmt"

func RecursivePower(nb int, power int) int {
	if power == 0{
		return 0
	}
	if power == 1 {
		return nb 
	}
	if power > 0 {
		return nb * RecursivePower(nb, power - 1)
	}
	return nb
}

// func main() {
// 	arg1 := 4
// 	arg2 := 3
// 	fmt.Println(RecursivePower(arg1, arg2))
// }