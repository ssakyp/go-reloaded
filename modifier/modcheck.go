package modifier

/*
(up)
(up, n)
(low)
(low, n)
(cap)
(cap, n)
(bin)
(hex)
*/

// Checks []rune for correspondance of modifiers and returns the modifier in string type and the specific
func CheckMod(sliceElem []rune) (string, int) {
	final := "0"
	spec := -1
	if len(sliceElem) >= 1 {
		i0 := sliceElem[0]

		switch len(sliceElem) {

		case 4:
			i1 := sliceElem[1]
			i2 := sliceElem[2]
			i3 := sliceElem[3]
			// (up) case
			if i0 == '(' && i1 == 'u' && i2 == 'p' && i3 == ')' {
				final = "(up)"
				// (up, case
			} else if i0 == '(' && i1 == 'u' && i2 == 'p' && i3 == ',' {
				final = "(up,"
				spec = 1
			}
		case 5:
			i1 := sliceElem[1]
			i2 := sliceElem[2]
			i3 := sliceElem[3]
			i4 := sliceElem[4]
			// (low) case
			if i0 == '(' && i1 == 'l' && i2 == 'o' && i3 == 'w' && i4 == ')' {
				final = "(low)"
			}
			// (low, case
			if i0 == '(' && i1 == 'l' && i2 == 'o' && i3 == 'w' && i4 == ',' {
				final = "(low,"
				spec = 1
			}
			// (cap) case
			if i0 == '(' && i1 == 'c' && i2 == 'a' && i3 == 'p' && i4 == ')' {
				final = "(cap)"
			}
			// (cap, case
			if i0 == '(' && i1 == 'c' && i2 == 'a' && i3 == 'p' && i4 == ',' {
				final = "(cap,"
				spec = 1
			}
			// (bin) case
			if i0 == '(' && i1 == 'b' && i2 == 'i' && i3 == 'n' && i4 == ')' {
				final = "(bin)"
			}
			// hex case
			if i0 == '(' && i1 == 'h' && i2 == 'e' && i3 == 'x' && i4 == ')' {
				final = "(hex)"
			}

		case 1:
			// article a case
			if i0 == 'a' {
				final = "a"
			}
			// article A case
			if i0 == 'A' {
				final = "A"
			}
		}
	}

	// spec is -1 if (cap)
	// spec is 1 if (cap, 2)
	return final, spec
}
