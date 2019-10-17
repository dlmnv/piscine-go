package piscine

/*import (
    "fmt"
)
*/
func UltimateDivMod(a *int, b *int) {
	var x int
	var z int
	x = *a / *b
	z = *a % *b
	*a = x
	*b = z
}

/*func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}*/
