package validators

func IsPunc(char rune) bool {
	if char == ',' || char == '.' || char == ':' || char == ';' || char == '!' || char == '?' || char == '#' {
		return true
	}
	return false
}

func IsMultiPunc(sliceElem []rune) bool {
	count := 0
	for _, char := range sliceElem {
		if IsPunc(char) {
			count++
		}
	}
	return count == len(sliceElem)
}

func IsApos(sliceElem []rune) int {
	n := -1
	if sliceElem[0] == '\'' {
		n = 0
	} else if sliceElem[len(sliceElem)-1] == '\'' {
		n = len(sliceElem) - 1
	}

	return n
}

func ContainsApos(slice [][]rune) bool {
	counter := 0

	for _, sliceElem := range slice {
		if IsApos(sliceElem) != -1 {
			counter++
		}
	}
	return counter%2 == 0
}

// checks if a slice contains in the beginnig or in the end the apostrophe
func ContApos(sliceElem []rune) bool {

	return IsApos(sliceElem) == -1 || IsApos(sliceElem) == 0
}

func IsPunct(sliceElem []rune) bool {
	puncList := ",.;?!:#"

	for _, punc := range puncList {
		if sliceElem[0] == punc {
			return true
		}
	}
	return false
}

func OnlyPunc(sliceElem []rune) bool {
	return len(sliceElem) == 1 && IsPunc(sliceElem[0])
}

// Functions for AposHandler
func HasSuffixApos(sliceElem []rune) bool {
	n := len(sliceElem)
	return sliceElem[n-1] == 39
}

func HasPrefixApos(sliceElem []rune) bool {
	
	return sliceElem[0] == 39 && len(sliceElem) > 1
}

func OnlyApos(sliceElem []rune) bool {
	return len(sliceElem) == 1 && sliceElem[0] == 39
}

func IsFullApos(sliceElem []rune) bool {
	return HasPrefixApos(sliceElem) && HasSuffixApos(sliceElem)
}

func IsPuncApos(sliceElem []rune) bool {
	n := 0
	for _, c := range sliceElem {
		if IsPunc(c) {
			n++
		}
		if c == '\'' {
			n++
		}
	}

	return n == len(sliceElem)
}