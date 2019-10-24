package piscine

func ConcatParams(args []string) string {
	var newstring string
	maslen := 0
	for range args {
		maslen++
	}
	for index := 0; index < maslen; index++ {
		newstring = newstring + args[index]
		if index != maslen-1 {
			newstring = newstring + "\n"
		}
	}
	return newstring
}
