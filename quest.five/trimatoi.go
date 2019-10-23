package main

import "fmt"

func TrimAtoi(s string) int {
	result := 0
	sign := 1

	for _, value := range s {
		if value >= '0' && value <= '9' {
			result = result*10 + int(value-48)
		} else if value == '-' && result == 0 {
			sign = -1
		}
	}

	return result * sign
}

func main() {
	s := "0-1234"


	n := TrimAtoi(s)


	fmt.Println(n)

}