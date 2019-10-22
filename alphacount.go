package main

// import "fmt"

func StrLen(str string) int {
	len := 0
	for range str {
		len++
	}
	return len
}

func AlphaCount(str string) int{
	k := []rune(str)
	x := StrLen(str)
	count := 0
	for i := 0; i < x; i++ {
		if k[i] < 91 || k[i] > 96 {
			if k[i] >= 'A' && k[i] <= 'z' {
				count++
			}
		}
	}
	return count
}

// func main() {
// 	str := "Hello 78 World!    4455 /"
// 	nb := AlphaCount(str)
// 	fmt.Println(nb)
// }
