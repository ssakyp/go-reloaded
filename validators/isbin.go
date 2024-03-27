package validators

func IsBin(slice []rune) bool {
	counter := 0

	for _, char := range slice {
		if char == '0' || char == '1' {
			counter++
		}
	}

	return counter == len(slice)
}
