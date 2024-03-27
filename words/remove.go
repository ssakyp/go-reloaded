package words

func RemoveElem(slice [][]rune, n int) [][]rune {
	var finalSlice [][]rune
	for i := 0; i <= len(slice)-1; i++ {
		elem := slice[i]
		if i != n {
			finalSlice = append(finalSlice, elem)
		}
	}

	return finalSlice
}

func TrimSuff(sliceElem []rune) []rune {
	var newSliceElem []rune
	for i, char := range sliceElem {
		if i != len(sliceElem)-1 {
			newSliceElem = append(newSliceElem, char)
		}
	}
	return newSliceElem
}

func TrimPrefix(sliceElem []rune) []rune {
	var newSliceElem []rune
	for i, char := range sliceElem {
		if i != 0 {
			newSliceElem = append(newSliceElem, char)
		}
	}
	return newSliceElem
}