package words

import "go-reloaded/validators"

// Capitalizes the sliceElement
func Capitalize(sliceElem []rune) []rune {
	for i, char := range sliceElem {
		if i == 0 && validators.IsAlpha(char) && char >= 'a' {
			sliceElem[i] = char - 32
		} else if i != 0 && validators.IsAlpha(char) && char < 'a' {
			sliceElem[i] = char + 32
		}
	}
	return sliceElem
}
