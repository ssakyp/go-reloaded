package validators

func IsVowel(char rune) bool {
	vowels := "aeiouhAEIOUH"
	for _, el := range vowels {
		if char == el {
			return true
		}
	}
	return false
}

func IsArtic(char rune) bool {
	return char == 'a' || char == 'A'
}
