package validators

func IsAlpha(char rune) bool {
	for i := 'a'; i <= 'z'; i++ {
		if i == char {
			return true
		}
	}
	for i := 'A'; i <= 'Z'; i++ {
		if i == char {
			return true
		}
	}
	return false
}
