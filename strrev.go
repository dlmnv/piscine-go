package main

// import "fmt"

func StrRev(s string) string {
	var x string
	z := 0
	for range s {
		z++
	}

	for a := z -1; a >= 0; a-- {
		x += string(s[a])
	}
	return x
}


// func main() {
// 	s := "Hello World!"
// 	s = StrRev(s)
// 	fmt.Println(s)
// }
