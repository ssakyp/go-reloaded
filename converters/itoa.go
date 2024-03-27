package converters

func Itoa(num int) []rune {
	// Handle the case when the number is 0
	if num == 0 {
		return []rune{'0'}
	}

	// Calculate the length of the number
	len := LenInt(num)

	// Initialize a slice to store the digits
	slice := make([]rune, len)

	// Fill the slice with digits of the number
	for i := len - 1; i >= 0; i-- {
		slice[i] = '0' + rune(num%10)
		num /= 10
	}

	return slice
}
