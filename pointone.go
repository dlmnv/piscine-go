package piscine

import ( "fmt" "github.com/01-edu/z01")


func PointOne(n *int) {
	*n = *n + 1
}

func main() {
	n := 0
	z01.PointOne(&n)
	fmt.Println()
}
