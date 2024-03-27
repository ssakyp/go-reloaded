package handlers

import (
	"go-reloaded/modifier"
	"go-reloaded/validators"
	"go-reloaded/words"
)

func CapHandlerCombined(i int, slice [][]rune, spec int) [][]rune {
	for k := 0; k < len(slice); k++ {
		if validators.IsCap(slice[k]) || validators.IsCapCom(slice[k]) {
			i = k
			break
		}
	}
	n := modifier.ModSpecFind(i, slice, spec)

	// case when the modifier is alone (up)
	if n >= 0 {
		return CapSpecHandler(i, slice, n)
	} else if n == -1 {
		return CapHandler(i, slice)
	}
	return slice

}

func CapHandler(i int, slice [][]rune) [][]rune {
	c := 0

	j := i - 1
	if j >= 0 && validators.IsWord(slice[j]) && !validators.IsMod(slice[j]) && !validators.OnlyPunc(slice[j]) && !validators.IsMultiPunc(slice[j]) && !validators.IsPuncApos(slice[j]) {
		if validators.IsPunc(slice[j][0]) && len(slice[j]) == 1 {
			slice = SinglePunc(slice, j)
			c = 1
			if len(slice) > 1 {

				slice[j-2] = words.Capitalize(slice[j-2])

			} else if len(slice) == 1 {

				slice[j-1] = words.Capitalize(slice[j-1])

			}
		} else {
			slice[j] = words.Capitalize(slice[j])
		}
	} else {
		for j >= 0 {
			if validators.IsWord(slice[j]) && !validators.OnlyPunc(slice[j]) && !validators.IsMultiPunc(slice[j]) && !validators.IsPuncApos(slice[j]) {
				slice[j] = words.Capitalize(slice[j])
				break
			}
			j--
		}
	}
	//5
	if c == 0 && i != 0 && !validators.ContainsParenth(slice[i-1]) {

		slice = words.RemoveElem(slice, i)

	} else {
		slice = words.RemoveElem(slice, i)
	}
	return slice

}

func CapSpecHandler(i int, slice [][]rune, spec int) [][]rune {
	// Calculate the number of elements to capitalize
	capCount := spec
	if capCount > i {
		capCount = i
	}
	finalBreak := false
	n := 0 // counter for case when before (low, n) there are no words
	m := 0 // counter for case when before (low, n) there are words

	if spec != 0 {
		// Uppers elements before index i
		for j := 0; j < capCount; j++ {
			if finalBreak {
				break
			}
			idx := i - j - 1
			if validators.IsWord(slice[idx]) && !validators.IsMod(slice[idx]) && !validators.OnlyPunc(slice[idx]) && !validators.IsMultiPunc(slice[idx]) && !validators.IsPuncApos(slice[idx]) {
				slice[idx] = words.Capitalize(slice[idx])
				m++
			} else { // case newline
				for idx >= 0 {
					if validators.IsWord(slice[idx]) && !validators.OnlyPunc(slice[idx]) && !validators.IsMultiPunc(slice[idx]) && !validators.IsPuncApos(slice[idx]) {
						slice[idx] = words.Capitalize(slice[idx])
						n++
					}
					if n+m == capCount {
						finalBreak = true
						break
					}
					idx--
				}
			}
		}
	}


	if i != 0 && !validators.IsMod(slice[i-1]) {
		// Remove the element at index i
		slice = words.RemoveElem(slice, i)
		slice = words.RemoveElem(slice, i)
	} else {
		slice = words.RemoveElem(slice, i)
		slice = words.RemoveElem(slice, i)
	}

	return slice
}
