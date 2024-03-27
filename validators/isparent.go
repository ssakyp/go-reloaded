package validators

func IsOpenParenth(char rune) bool {
	return char == '('
}

func IsClosedParenth(char rune) bool {
	return char == ')'
}

func ContainsParenth(sliceElem []rune) bool {
	for _, char := range sliceElem {
		if IsOpenParenth(char) || IsClosedParenth(char) {
			return true
		}
	}
	return false
}
