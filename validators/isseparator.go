package validators

func IsSpace(char rune) bool {
	return char == ' '
}

func IsNewLine(char rune) bool {
	return char == '\n'
}

func IsNewLineElem(sliceElem []rune) bool {
	return IsNewLine(sliceElem[0]) && len(sliceElem) == 1
}