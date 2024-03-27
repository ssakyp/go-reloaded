package handlers

import "go-reloaded/validators"

func ArticleHandler(slice [][]rune) [][]rune {
	for i := 0; i < len(slice)-1; i++ {
		if validators.IsArtic(slice[i][0]) && len(slice[i]) == 1 {
			if validators.IsVowel(slice[i+1][0]) && len(slice[i+1]) > 1 {
				slice[i] = append(slice[i], 'n')
			}
		}

	}
	return slice
}
