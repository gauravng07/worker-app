package worker

import (
	"scb-recipe-app/internal/model"
	"sort"
)

var specialWords = []string{"Potato", "Veggie", "Mushroom"}

func Start(ch chan model.RecipeStats, result chan model.Response, postCode string, deliveryStartTime string, deliveryEndTime string, recipeName string) {
	response := model.Response{}
	uniqueRecipes := make(map[string]int, 0)
	uniquePostCode := make(map[string]int, 0)
	countPerRecipe := make([]model.RecipeType, 0)
	var busiestPostCode model.BusiestPostCode
	var countPerPostCodeAndTime model.CountPerPostCodeAndTime
	uniqueRecipesCount := 0
	var defaultPostCodeDeliveryCount int
	var matchByName []string

	for stat := range ch {
		if _, ok := uniqueRecipes[stat.Recipe]; !ok {
			uniqueRecipesCount = uniqueRecipesCount + 1
			recipeType := model.RecipeType{
				Recipe: stat.Recipe,
				Count: 1,
			}
			countPerRecipe = append(countPerRecipe, recipeType)
			uniqueRecipes[stat.Recipe] = uniqueRecipesCount
			if recipeHasSpamWords(stat.Recipe, specialWords) {
				matchByName = append(matchByName, stat.Recipe)
			}
		} else {
			countPerRecipe = updateRecipeCount(countPerRecipe, stat.Recipe)
		}

		if _, ok := uniquePostCode[stat.PostCode]; !ok {
			uniquePostCode[stat.PostCode] = 1
		} else {
			uniquePostCode[stat.PostCode] = uniquePostCode[stat.PostCode] + 1
		}

		if stat.PostCode == postCode {
			start, end := parseDelivery(stat.Delivery);
			if  postCodeDeliveryTime(start, end, deliveryStartTime, deliveryEndTime) {
				defaultPostCodeDeliveryCount = defaultPostCodeDeliveryCount + 1
			}
		}
	}

	if _, ok := uniquePostCode[postCode]; ok {
		countPerPostCodeAndTime.PostCode = postCode
		countPerPostCodeAndTime.From = deliveryStartTime
		countPerPostCodeAndTime.To = deliveryEndTime
		countPerPostCodeAndTime.DeliveryCount = defaultPostCodeDeliveryCount
	}

	sort.Slice(matchByName, func(i, j int) bool {
		return i < j
	})

	sort.Slice(countPerRecipe, func(i, j int) bool {
		return countPerRecipe[i].Recipe < countPerRecipe[j].Recipe
	})

	response.MatchByName = matchByName
	response.UniqueRecipeCount = uniqueRecipesCount
	response.CountPerRecipe = countPerRecipe
	response.CountPerPostCodeAndTime = countPerPostCodeAndTime

	postCode, count := rankMapStringInt(uniquePostCode)
	busiestPostCode.PostCode = postCode
	busiestPostCode.DeliveryCount = count
	response.BusiestPostCode = busiestPostCode;
	result <- response
}

func StatByRecipeName(ch chan model.RecipeStats, recipeName string)  model.RecipeType {
	var recipeType model.RecipeType
	var recipesCount int
	for stat := range ch {
		if stat.Recipe == recipeName {
			recipesCount = recipesCount + 1
		}
	}
	recipeType.Recipe = recipeName
	recipeType.Count = recipesCount
	return recipeType
}