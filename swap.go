package piscine

func Swap(a *int, b *int) {
	x := *b
	z := *a
	*a = x
	*b = z
}
