package converters

// Function to convert a binary number represented as a slice of runes to its decimal equivalent
func BinToDecimal(bin []rune) int {
	decimal := 0
	power := 1 // 2^0 = 1

	// Iterate over each character of the binary string from right to left
	for i := len(bin) - 1; i >= 0; i-- {
		char := bin[i]

		// Convert binary digit to its decimal value
		digit := int(char - '0')

		// Multiply the decimal value of the digit by the corresponding power of 2
		decimal += digit * power

		// Update the power for the next digit
		power *= 2
	}

	return decimal
}
