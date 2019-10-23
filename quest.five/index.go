package piscine

func Index(s string, toFind string) int {
	srune := []rune(s)
	frune := []rune(toFind)
	lens := 0
	lenf := 0
	for index := range frune {
		index = index
		lenf++
	}
	if lenf == 0 {
		return 0
	}
	for index := range srune {
		index = index
		lens++
	}
	for index, letter := range srune {
		if letter == frune[0] && lens >= lenf+index-1 {
			m := 1
			for i := 1; i < lenf; i++ {
				if frune[i] == srune[index+i] {
					m++
				}
			}
			if m == lenf {
				return index
			}
		}
	}
	return -1
}
