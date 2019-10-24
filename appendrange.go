package main

import "fmt"

func AppendRange(min, max int) []int {
	var arr []int
	for k := min; k < max; k++ {
		arr = append(arr, k)
	}
	return arr
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
