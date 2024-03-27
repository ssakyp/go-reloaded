package words

import "go-reloaded/validators"

// Lowers all chars
func ToLower(sliceElem []rune) []rune {
	for i, char := range sliceElem {
		if validators.IsAlpha(char) && char <= 'Z' {
			sliceElem[i] = char + 32
		}
	}
	return sliceElem
}
