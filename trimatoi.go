package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1

	for _, value := range s {
		if value >= '0' && value <= '9' {
			result = result*10 + int(value-48)
		} else if value == '-' && result == 0 {
			sign = -1
		} else if value == '0' && result == 0 {
			return
		}
	}

	return result * sign
}
