package handlers

import (
	"go-reloaded/converters"
	"go-reloaded/validators"
	"go-reloaded/words"
)

func HexHandler(i int, slice [][]rune) [][]rune {

	for k := 0; k < len(slice); k++ {
		if validators.IsHexMod(slice[k]) {
			i = k
		}
	}
	j := i - 1
	isRemoved := false

	if i != 0 && validators.IsHex(slice[i-1]) {
		slice[i-1] = converters.Itoa(converters.HexToDecimal(slice[i-1]))
		slice = words.RemoveElem(slice, i)
		isRemoved = true
	} else if i > 0 && validators.IsNewLineElem(slice[i-1]) {
		for j >= 0 {
			if !validators.IsNewLineElem(slice[j]) {
				if validators.IsHex(slice[j]) {
					slice[j] = converters.Itoa(converters.HexToDecimal(slice[j]))
					slice = words.RemoveElem(slice, i)
					isRemoved = true
					break
				} else {
					break
				}
			}
			j--
		}
	}

	if !isRemoved {
		slice = words.RemoveElem(slice, i)
	}
	return slice
}
