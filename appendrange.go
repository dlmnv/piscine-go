package piscine

func AppendRange(min, max int) []int {
	var arr []int
	for k := min; k < max; k++ {
		arr = append(arr, k)
	}
	return arr
}
