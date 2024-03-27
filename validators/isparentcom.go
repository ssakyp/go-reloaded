package validators

func IsLastParent(sliceElem []rune) bool {
	return len(sliceElem)-1 == ')'
}

func IsFirstComma(sliceElem []rune) bool {
	return sliceElem[0] == ','
}
