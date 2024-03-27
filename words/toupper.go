package words

import "go-reloaded/validators"

// Uppers all chars
func ToUpper(sliceElem []rune) []rune {
	for i, char := range sliceElem {
		if validators.IsAlpha(char) && char >= 'a' {
			sliceElem[i] = char - 32
		}
	}
	return sliceElem
}
