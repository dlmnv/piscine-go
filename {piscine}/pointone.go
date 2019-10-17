package piscine

func PointOne(a *int) {
	*a = *a + 1
}

func main() {
	n := 0
	z01.PointOne(&n)
	fmt.Println(n)
}
