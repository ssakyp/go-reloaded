package words

import (
	"go-reloaded/validators"
)

func Split(str string) [][]rune {
	var slice [][]rune
	word := ""
	// isWord := false

	for _, char := range str {
		if char == ' ' && word != "" {
			slice = append(slice, []rune(word))
			word = ""
		} else if validators.IsWordRune(char) {
			word += string(char)
		} else if char == '\n' && word != "" {
			slice = append(slice, []rune(word))
			word = ""
			slice = append(slice, []rune{'\n'})
		} else if char == '\n' {
			slice = append(slice, []rune{'\n'})
		}
	}
	if word != "" {
		slice = append(slice, []rune(word))
		word = ""
	}
	return slice
}
