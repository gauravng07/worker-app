package worker

import (
	"hello-fresh-app/internal/model"
	"strings"
)

func recipeHasSpamWords(text string, specialWords []string) bool {
	removePunctuation := func(r rune) rune {
		if strings.ContainsRune(".,:;", r) {
			return -1
		} else {
			return r
		}
	}
	text = strings.Map(removePunctuation, text)
	textArr := strings.Split(text, ` `)
	for _, word := range textArr {
		for _, sapmWord := range specialWords {
			if word == sapmWord {
				return true
			}
		}
	}
	return false
}


func updateRecipeCount(countPerRecipe []model.RecipeType, recipe string) []model.RecipeType {
	for i := 0; i < len(countPerRecipe) ; i++ {
		if countPerRecipe[i].Recipe == recipe {
			countPerRecipe[i].Count = countPerRecipe[i].Count + 1
			break
		}
	}
	return countPerRecipe
}