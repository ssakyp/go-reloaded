package converters

func SliceRuneToSliceStr(slice [][]rune) []string {
	var final []string
	word := ""

	for _, el := range slice {
		for _, char := range el {
			word += string(char)
		}
		final = append(final, word)
		word = ""
	}
	return final
}
