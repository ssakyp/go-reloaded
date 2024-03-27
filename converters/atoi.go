package converters

import "go-reloaded/validators"

func Atoi(slice []rune) int {
	result := 0
	if validators.IsNumber(slice) {
		for _, digit := range slice {
			result = result*10 + int(digit-'0')
		}
	}

	return result
}
