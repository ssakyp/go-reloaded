package validators

func IsHex(sliceElem []rune) bool {
	counter := 0
	hexChars := "0123456789abcdefABCDEF"

	for _, char := range sliceElem {
		for _, hexChar := range hexChars {
			if char == hexChar {
				counter++
			}
		}
	}

	return counter == len(sliceElem)
}
