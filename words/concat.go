package words

func AddElem(slice []rune, sliceToAdd []rune) []rune {
	slice = append(slice, sliceToAdd...)
	return slice
}

func Concat(slice [][]rune) string {
	
	final := ""
	for i, elem := range slice {
		if elem[0] == '\n' && len(elem) == 1 {
			final += "\n"
			continue
		}

		if i != 0 && slice[i-1][0] != '\n' {
			final += " "
		}

		for _, c := range elem {
			if c != '#' {
				final += string(c)
			}
		}
		// if elem[0] != '#' {
		// 	final += string(elem)
		// }

	}

	return final
}

func AddSuffix(sliceElem []rune, r rune) []rune {
	sliceElem = append(sliceElem, r)
	return sliceElem
}

func AddPrefix(sliceElem []rune, r rune) []rune {
	var final []rune = []rune{r}
	final = append(final, sliceElem...)
	return final
}
