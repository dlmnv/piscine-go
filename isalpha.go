package main

import "fmt"

func IsAlpha(str string) bool {
	slen := 0
	x := []rune(str)
	for range str {
		slen++
	}
	for i := 0; i < slen; i++ {
		if x[i] < '0' || (x[i] > '9' && x[i] < 'A') || (x[i] > 'Z' && x[i] < 'a') || x[i] > 'z' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}
