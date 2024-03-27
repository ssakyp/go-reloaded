package handlers

import (
	"go-reloaded/modifier"
	"go-reloaded/validators"
	"go-reloaded/words"
)

func LowHandler(i int, slice [][]rune) [][]rune {
	l := 0
	j := i - 1
	if j >= 0 && validators.IsWord(slice[j]) && !validators.IsMod(slice[j]) && !validators.OnlyPunc(slice[j]) && !validators.IsMultiPunc(slice[j]) && !validators.IsPuncApos(slice[j]) {
		if validators.IsPunc(slice[j][0]) && len(slice[j]) == 1 {
			slice = SinglePunc(slice, j)
			l = 1
			if len(slice) > 1 {
				slice[j-2] = words.ToLower(slice[j-2])
			} else if len(slice) == 1 {
				slice[j-1] = words.ToLower(slice[j-1])
			}
		} else {
			slice[j] = words.ToLower(slice[j])
		}
	} else {
		for j >= 0 {
			if validators.IsWord(slice[j]) && !validators.OnlyPunc(slice[j]) && !validators.IsMultiPunc(slice[j]) && !validators.IsPuncApos(slice[j]) {
				slice[j] = words.ToLower(slice[j])
				break
			}
			j--
		}
	}

	if l == 0 && i != 0 && !validators.IsMod(slice[i-1]) {
		slice = words.RemoveElem(slice, i)
	} else {
		slice = words.RemoveElem(slice, i)
	}

	return slice
}
func LowSpecHandler(i int, slice [][]rune, spec int) [][]rune {
	// Calculate the number of elements to lower
	lowCount := spec
	if lowCount > i {
		lowCount = i
	}

	finalBreak := false
	n := 0 // counter for case when before (low, n) there are no words
	m := 0 // counter for case when before (low, n) there are words

	if spec != 0 {
		// Lowers elements before index i
		for j := 0; j < lowCount; j++ {
			if finalBreak {
				break
			}
			idx := i - j - 1

			if validators.IsWord(slice[idx]) && !validators.IsMod(slice[idx]) && !validators.OnlyPunc(slice[idx]) && !validators.IsMultiPunc(slice[idx]) && !validators.IsPuncApos(slice[idx]){
				slice[idx] = words.ToLower(slice[idx])
				m++
			} else { // case newline
				for idx >= 0 {
					if validators.IsWord(slice[idx]) && !validators.OnlyPunc(slice[idx]) && !validators.IsMultiPunc(slice[idx]) && !validators.IsPuncApos(slice[idx]) {
						slice[idx] = words.ToLower(slice[idx])
						n++
					}
					if n+m == lowCount {
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

func LowHandlerCombined(i int, slice [][]rune, spec int) [][]rune {

	for k := 0; k < len(slice); k++ {
		if validators.IsLow(slice[k]) || validators.IsLowCom(slice[k]) {
			i = k
			break
		}
	}

	n := modifier.ModSpecFind(i, slice, spec)

	// case when the modifier is alone (up)
	if n >= 0 {
		return LowSpecHandler(i, slice, n)
	} else if n == -1 {
		return LowHandler(i, slice)
	}
	return slice
}
