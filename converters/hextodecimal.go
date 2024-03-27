package converters

func HexToDecimal(hex []rune) int {
	decimal := 0
	power := 1 // 16^0 = 1

	// Iterate over each character of the hexadecimal string from right to left
	for i := len(hex) - 1; i >= 0; i-- {
		char := hex[i]

		// Convert hexadecimal digit to its decimal value
		var digit int
		switch {
		case '0' <= char && char <= '9':
			digit = int(char - '0')
		case 'A' <= char && char <= 'F':
			digit = int(char - 'A' + 10)
		case 'a' <= char && char <= 'f':
			digit = int(char - 'a' + 10)
		default:
			return -1 // Invalid hexadecimal character
		}

		// Multiply the decimal value of the digit by the corresponding power of 16
		decimal += digit * power

		// Update the power for the next digit
		power *= 16
	}

	return decimal
}
