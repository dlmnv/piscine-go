package main

import "fmt"

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	slice := make([]int, max-min)
	for j, i := 0, min; i < max; j, i = j+1, i+1 {
		slice[j] = i
	}
	return slice
}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
