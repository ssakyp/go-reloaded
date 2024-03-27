package handlers

import (
	"go-reloaded/validators"
	"go-reloaded/words"
)

func PuncHandler(slice [][]rune) [][]rune {

	for i := len(slice) - 1; i >= 0; i-- {
		if validators.IsPunct(slice[i]) {
			//handles the cases ... ?!
			if validators.IsMultiPunc(slice[i]) {
				if i > 0 {
					slice[i-1] = words.AddElem(slice[i-1], slice[i])
					slice = words.RemoveElem(slice, i)
				}

			} else if i > 0 && len(slice[i]) > 1 && !validators.IsMultiPunc(slice[i]) {
				for validators.IsPuncWord(slice[i]) {
					slice[i-1] = append(slice[i-1], slice[i][0])
					slice[i] = words.TrimPrefix(slice[i])
				}

			} else if i > 0 && len(slice[i]) == 1 {
				slice[i-1] = append(slice[i-1], slice[i][0])
				slice = words.RemoveElem(slice, i)
			}
		}

	}

	return slice
}

func SinglePunc(slice [][]rune, j int) [][]rune {
	for i := 0; i < len(slice); i++ {
		if i == j {
			slice[i-1] = append(slice[i-1], slice[i][0])
			slice = words.RemoveElem(slice, i)
			//removes the operator
			slice = words.RemoveElem(slice, i)
		}
	}
	return slice
}
