package model

type RecipeStats struct {
	PostCode	string		`json:"postcode"`
	Recipe		string		`json:"recipe"`
	Delivery 	string 		`json:"delivery"`
}

type RecipeType struct {
	Recipe		string 		`json:"recipe"`
	Count		int			`json:"count"`
}

type BusiestPostCode struct {
	PostCode		string		`json:"postcode"`
	DeliveryCount	int			`json:"delivery_count"`
}

type CountPerPostCodeAndTime struct {
	PostCode		string		`json:"postcode"`
	From			string		`json:"from"`
	To				string		`json:"to"`
	DeliveryCount	int			`json:"delivery_count"`
}

type Response struct {
	UniqueRecipeCount 			int							`json:"unique_recipe_count"`
	CountPerRecipe				[]RecipeType 				`json:"count_per_recipe"`
	BusiestPostCode 			BusiestPostCode				`json:"busiest_post_code"`
	CountPerPostCodeAndTime 	CountPerPostCodeAndTime 	`json:"count_per_post_code_and_time"`
	MatchByName					[]string					`json:"match_by_name"`
}

type Stats struct {
	RecipesName string
	PostCode	string
	DeliveryTime	string
}
