package validators

func IsWordRune(char rune) bool {
	if char >= '!' && char <= '~' {
		return true
	}
	
	return false
}

func IsPuncWord(elem []rune) bool {
	n := 0
	for _, char := range elem {
		if IsWordRune(char) {
			n++
		}
	}
	return n == len(elem) && IsPunc(elem[0])
}

func IsWordPunc(elem []rune) bool {
	containsAlpha := false
	containsPunc := false
	for _, char := range elem {
		if IsWordRune(char) {
			containsAlpha = true
		}
		if IsPunc(char) {
			containsPunc = true
		}
	}
	return containsAlpha && containsPunc
}

func IsWord(elem []rune) bool {
	n := 0
	for _, char := range elem {
		if IsWordRune(char) {
			n++
		}
	}

	return n == len(elem)
}