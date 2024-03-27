package validators

func IsNum(char rune) bool {
	for i := '0'; i <= '9'; i++ {
		if i == char {
			return true
		}
	}
	return false
}

// 23) => len = 3
// 23 => len = 2
// This function checks if the elements before the last one are numbers
func IsNumUntilLast(sliceElem []rune) bool {
	counter := 0

	for i := 0; i < len(sliceElem)-1; i++ {
		if IsNum(sliceElem[i]) {
			counter++
		}
	}

	return counter == len(sliceElem)-1
}

func IsNumber(sliceElem []rune) bool {
	counter := 0

	for _, char := range sliceElem {
		if IsNum(char) {
			counter++
		}
	}
	return counter == len(sliceElem)
}

func IsWordNum(str string) bool {
	for _, char := range str {
		if IsAlpha(char) || IsNum(char) || IsPunc(char) {
			return false
		}
	}
	return true
}
