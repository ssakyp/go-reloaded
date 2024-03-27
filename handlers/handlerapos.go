package handlers

import (
	"go-reloaded/validators"
	"go-reloaded/words"
)

func ApostropheHandler(slice [][]rune) [][]rune {
	isApos := false

	counter := 0 // for identifying even number of apostrophies

	for _, elem := range slice {
		if validators.OnlyApos(elem) || validators.HasSuffixApos(elem) || validators.HasPrefixApos(elem) {
			counter++
		}
	}

	if counter%2 != 0 {
		return slice
	}

	for i := 0; i < len(slice); i++ {
		elem := slice[i]

		if i > 0 && !isApos && validators.HasPrefixApos(slice[i-1]) && !validators.IsFullApos(slice[i-1]) {
			isApos = true
		}

		if !isApos {
			if validators.OnlyApos(elem) {

				if i < len(slice)-1 && !validators.IsNewLineElem(slice[i+1]) {
					slice[i] = []rune{'#'}

					slice[i+1] = words.AddPrefix(slice[i+1], '\'')

					isApos = true
				} else {
					continue
				}

				// i++
			} else if validators.HasSuffixApos(elem) && !validators.HasPrefixApos(elem) {

				if !validators.IsNewLineElem(slice[i+1]) {
					slice[i+1] = words.AddPrefix(slice[i+1], '\'')

					slice[i] = words.TrimSuff(slice[i])

					isApos = true
				} else {
					continue
				}

				// i++
			}

			if i < len(slice)-1 && (validators.HasSuffixApos(elem) || validators.OnlyApos(elem)) {
				i++
			}

		} else {
			if validators.OnlyApos(elem) {
				if validators.OnlyPunc(slice[i-1]) {
					if i > 2 {
						slice[i-2] = words.AddSuffix(slice[i-2], slice[i-1][0])
						slice[i-1] = []rune{'#'}
					}
					slice[i-2] = words.AddSuffix(slice[i-2], '\'')

					slice[i] = []rune{'#'}

					isApos = false
					continue
				}
				slice[i-1] = words.AddSuffix(slice[i-1], '\'')

				slice[i] = []rune{'#'}

				isApos = false
			} else if validators.HasPrefixApos(elem) {

				slice[i-1] = words.AddSuffix(slice[i-1], '\'')

				slice[i] = words.TrimPrefix(slice[i])

				isApos = false
			}
		}

	}

	return slice
}
