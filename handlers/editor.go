package handlers

import (
	"go-reloaded/modifier"
	"go-reloaded/words"
)

func Editor(slice [][]rune) string {
	edited := ""
	for i, sliceElem := range slice {
		mod, spec := modifier.CheckMod(sliceElem)
		switch mod {
		case "(up)":
			slice = UpHandlerCombined(i, slice, spec)
		case "(up,":
			slice = UpHandlerCombined(i, slice, spec)
		case "(low)":
			slice = LowHandlerCombined(i, slice, spec)
		case "(low,":
			slice = LowHandlerCombined(i, slice, spec)
		case "(cap)":
			slice = CapHandlerCombined(i, slice, spec)
		case "(cap,":
			slice = CapHandlerCombined(i, slice, spec)
		case "(hex)":
			slice = HexHandler(i, slice)
		case "(bin)":
			slice = BinHandler(i, slice)
		}
	}

	slice = ApostropheHandler(slice)
	slice = PuncHandler(slice)

	slice = ArticleHandler(slice)
	edited = words.Concat(slice)
	editedSlice := words.Split(edited)
	edited = words.Concat(editedSlice)
	return edited
}
