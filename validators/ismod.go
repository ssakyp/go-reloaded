package validators

func IsCap(sliceElem []rune) bool {
	c := sliceElem
	return c[0] == '(' && c[1] == 'c' && c[2] == 'a' && c[3] == 'p' && c[4] == ')'
}
func IsCapCom(sliceElem []rune) bool {
	c := sliceElem
	return c[0] == '(' && c[1] == 'c' && c[2] == 'a' && c[3] == 'p' && c[4] == ','
}
func IsUp(sliceElem []rune) bool {
	c := sliceElem
	return c[0] == '(' && c[1] == 'u' && c[2] == 'p' && c[3] == ')'
}

func IsUpCom(sliceElem []rune) bool {
	c := sliceElem
	return c[0] == '(' && c[1] == 'u' && c[2] == 'p' && c[3] == ','
}

func IsLow(sliceElem []rune) bool {
	c := sliceElem
	return c[0] == '(' && c[1] == 'l' && c[2] == 'o' && c[3] == 'w' && c[4] == ')'
}

func IsLowCom(sliceElem []rune) bool {
	c := sliceElem
	return c[0] == '(' && c[1] == 'l' && c[2] == 'o' && c[3] == 'w' && c[4] == ','
}

func IsHexMod(sliceElem []rune) bool {
	c := sliceElem
	return c[0] == '(' && c[1] == 'h' && c[2] == 'e' && c[3] == 'x' && c[4] == ')'
}

func IsBinMod(sliceElem []rune) bool {
	c := sliceElem
	return c[0] == '(' && c[1] == 'b' && c[2] == 'i' && c[3] == 'n' && c[4] == ')'
}

func IsMod(sliceElem []rune) bool {
	if len(sliceElem) == 5 && (IsCap(sliceElem) || IsLow(sliceElem) || IsHexMod(sliceElem) || IsBinMod(sliceElem)) {
		return true
	} else if len(sliceElem) == 4 && IsUp(sliceElem) {
		return true
	}
	return false
}
