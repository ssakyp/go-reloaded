package modifier

import (
	"go-reloaded/converters"
	"go-reloaded/validators"
	"go-reloaded/words"
)

func ModSpecFind(index int, slice [][]rune, spec int) int {

	n := 0
	var nrune []rune

	if spec == -1 {
		n = -1
	} else if index < len(slice)-1 && validators.IsNumUntilLast(slice[index+1]) && spec != 0 {
		nrune = words.TrimSuff(slice[index+1])
		n = converters.Atoi(nrune)
	} else if index < len(slice)-1 && !validators.IsNumUntilLast(slice[index+1]) {
		n = -2
	}

	//Returns -1 when it is only (up), (low), (cap),...
	return n
}
