package converters

func LenInt(num int) int {
	// Handle the case when the number is 0
	if num == 0 {
		return 1
	}

	len := 0

	for num != 0 {
		num /= 10
		len++
	}

	return len
}
